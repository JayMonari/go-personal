package main

import (
	"context"
	"log"
	"net"

	pb "example.com/go-usrmgmt-grpc/usrmgmt"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
)

const (
	port = ":9001"
)

type UserManagementServer struct {
	conn *pgx.Conn
	pb.UnimplementedUserManagementServer
}

func (svc *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, svc)
	log.Println("server listening on", lis.Addr())
	return s.Serve(lis)
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Recieved: %q\n", in.GetName())

	createSQL := `
	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		name TEXT,
		age INT
	);`
	_, err := s.conn.Exec(ctx, createSQL)
	if err != nil {
		log.Fatal(err)
	}
	u := &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
	}
	tx, err := s.conn.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback(ctx)
	defer tx.Conn().Close(ctx)

	if _, err = tx.Exec(ctx, "INSERT INTO users(name, age) VALUES ($1, $2)", u.Name, u.Age); err != nil {
		log.Fatal(err)
	}
	tx.Commit(ctx)
	return u, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	ul := pb.UserList{}
	rows, err := s.conn.Query(ctx, "SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		u := pb.User{}
		err = rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			return nil, err
		}
		ul.Users = append(ul.Users, &u)
	}
	return &ul, nil
}

const dbURL = "postgres://postgres:password@localhost:5432/postgres"

func main() {
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	if err := (&UserManagementServer{conn: conn}).Run(); err != nil {
		log.Fatal(err)
	}
}
