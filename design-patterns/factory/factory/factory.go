package factory

import "factory/connection"

type DBType uint8

const (
	MySQL DBType = iota
	Postgres
)

func Factory(t DBType) connection.DBConnection {
	switch t {
	case MySQL:
		return &connection.MySQL{}
	case Postgres:
		return &connection.Postgres{}
	default:
		panic("How did you get here?")
	}
}
