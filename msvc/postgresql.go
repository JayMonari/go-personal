package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgeSQL struct {
	pool *pgxpool.Pool
}

func NewPostgreSQL() (*PostgeSQL, error) {
	p, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &PostgeSQL{pool: p}, nil
}

func (p *PostgeSQL) Close() { p.pool.Close() }

func (p *PostgeSQL) FindByNConst(nconst string) (res Name, err error) {
	query := `SELECT nconst, primary_name, birth_year, death_year FROM "names" WHERE nconst = $1`

	if err := p.pool.QueryRow(context.Background(), query, nconst).
		Scan(&res.NConst, &res.Name, &res.BirthYear, &res.DeathYear); err != nil {
		return Name{}, err
	}
	return res, nil
}
