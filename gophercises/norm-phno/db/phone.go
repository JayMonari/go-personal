package db

import "database/sql"

// Phone represents the phone_numbers table in the DB
type Phone struct {
	ID     int
	Number string
}

func (db *DB) DeletePhone(id int) error {
	q := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.db.Exec(q, id)
	return err
}

func (db *DB) UpdatePhone(p *Phone) error {
	q := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.db.Exec(q, p.ID, p.Number)
	return err
}

func (db *DB) FindPhone(number string) (*Phone, error) {
	var p Phone
	row := db.db.QueryRow("SELECT * FROM phone_numbers WHERE id=$1", number)
	err := row.Scan(&p.ID, &p.Number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (db *DB) AllPhones() ([]Phone, error) {
	rows, err := db.db.Query("SELECT id, value FROM phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pp []Phone
	for rows.Next() {
		var p Phone
		if err := rows.Scan(&p.ID, &p.Number); err != nil {
			return nil, err
		}
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return pp, nil
}

type DB struct {
	db *sql.DB
}

func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Seed() error {
	data := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"(123)456-7892",
	}
	for _, v := range data {
		if _, err := insertPhone(db.db, v); err != nil {
			return err
		}
	}
	return nil
}

func insertPhone(db *sql.DB, number string) (int, error) {
	q := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(q, number).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func Migrate(driverName, dataSourceName string) error {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}
	err = createTable(db)
	if err != nil {
		return err
	}
	return db.Close()
}

func createTable(db *sql.DB) error {
	q := `
  CREATE TABLE IF NOT EXISTS phone_numbers (
    id SERIAL,
    value VARCHAR(255)
  )`
	_, err := db.Exec(q)
	return err
}

func Reset(driverName, dataSourceName, dbName string) error {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}
	err = resetDB(db, dbName)
	if err != nil {
		return err
	}
	return db.Close()
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	return err
}
