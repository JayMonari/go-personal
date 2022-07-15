package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/jackc/pgx/v4"
)

func main() {
	amount, size := parseFlags()
	setUpFaker()

	//-

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Connecting to db %v", err)
	}

	names, errs := make(chan name), make(chan error)
	go func() {
		defer func() {
			close(names)
			close(errs)
		}()

		t := time.NewTicker(5 * time.Second)
		for ; amount > 0; amount-- {
			select {
			case <-t.C:
				fmt.Printf("%.2f%% out of 100%% \n", 100-float32(amount)/500_000*100)
			default:
			}
			n := name{}
			if err := faker.FakeData(&n); err != nil {
				errs <- err
				return
			}
			names <- n
		}
	}()

	wait := make(chan struct{})
	go func() {
		defer close(wait)

		batcher := batcher{conn: conn, size: size}
		errs := batcher.Copy(context.Background(), names)
		if err := <-errs; err != nil {
			fmt.Printf("Error copying %v\n", err)
		}
	}()
	<-wait
	fmt.Println("Done")
}

//-

type name struct {
	NConst             string
	PrimaryName        string   `faker:"name"`
	BirthYear          string   `faker:"yearString"`
	DeathYear          string   `faker:"yearString"`
	PrimaryProfessions []string `faker:"len=5"`
	KnownForTitles     []string `faker:"len=7"`
}

func (n name) Values() []any {
	v := make([]any, 6)
	v[0] = n.NConst
	v[1] = n.PrimaryName
	v[2] = n.BirthYear
	v[3] = n.DeathYear
	v[4] = n.PrimaryProfessions
	v[5] = n.KnownForTitles
	return v
}

//-

type copyFromSource struct {
	errors  <-chan error
	names   <-chan name
	err     error
	closed  bool
	current name
}

func (c *copyFromSource) Err() error { return c.err }

func (c *copyFromSource) Next() bool {
	if c.closed {
		return false
	}
	var open bool
	select {
	case c.current, open = <-c.names:
	case c.err = <-c.errors:
	}
	if !open {
		c.closed = true
		return false
	}
	if c.err != nil {
		return false
	}
	return true
}

func (c *copyFromSource) Values() ([]any, error) {
	return c.current.Values(), nil
}

//-

type copyFromSourceMediator struct {
	names  chan name
	errors chan error
	copier *copyFromSource
}

func newCopyFromSourceMediator(conn *pgx.Conn) (*copyFromSourceMediator, <-chan error) {
	errors := make(chan error)
	names := make(chan name)
	copier := copyFromSource{errors: errors, names: names}
	res := copyFromSourceMediator{
		names:  names,
		errors: errors,
		copier: &copier,
	}

	outErrs := make(chan error)
	go func() {
		defer close(outErrs)

		_, err := conn.CopyFrom(context.Background(),
			pgx.Identifier{"names"},
			[]string{
				"nconst",
				"primary_name",
				"birth_year",
				"death_year",
				"primary_professions",
				"known_for_titles",
			}, &copier)

		outErrs <- err
	}()

	return &res, outErrs
}

func (c *copyFromSourceMediator) Batch(n name) { c.names <- n }

func (c *copyFromSourceMediator) Err(err error) { c.errors <- err }

func (c *copyFromSourceMediator) CopyAll() {
	close(c.names)
	close(c.errors)
}

//-

type batcher struct {
	conn *pgx.Conn
	size int
}

func (b *batcher) Copy(ctx context.Context, names <-chan name) <-chan error {
	outErrs := make(chan error)
	go func() {
		mediator, errs := newCopyFromSourceMediator(b.conn)
		copyAll := func(m *copyFromSourceMediator, c <-chan error) error {
			m.CopyAll()
			return <-c
		}
		defer func() {
			if err := copyAll(mediator, errs); err != nil {
				outErrs <- err
			}
			close(outErrs)
		}()

		var index int
		for {
			select {
			case name, open := <-names:
				if !open {
					return
				}
				mediator.Batch(name)

				if index++; index == b.size {
					if err := copyAll(mediator, errs); err != nil {
						outErrs <- err
					}
					mediator, errs = newCopyFromSourceMediator(b.conn)
					index = 0
				}
			case err := <-errs:
				outErrs <- err
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					mediator.Err(err)
					outErrs <- err
				}
			}
		}
	}()

	return outErrs
}

//-

func parseFlags() (int, int) {
	var amount, size int

	flag.IntVar(&size, "size", 100_000, "batch size")
	flag.IntVar(&amount, "amount", 500_000, "amount of fakes to generate")
	flag.Parse()

	if size <= 0 {
		size = 100_000
	}
	if amount <= 0 {
		amount = 500_000
	}
	return amount, size
}

func setUpFaker() {
	_ = faker.AddProvider("name", func(_ reflect.Value) (any, error) {
		return faker.Name(), nil
	})

	_ = faker.AddProvider("yearString", func(_ reflect.Value) (any, error) {
		return faker.YearString(), nil
	})
}
