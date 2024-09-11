package postgre

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host = "db"
	port = 5432
)

var pool *sql.DB

func InitDB(pathToSchema string) error {
	var err error
	if pool == nil {
		err = initConnection()
		if err != nil {
			pool = nil
			return errors.Join(fmt.Errorf("couldn't connect to db"), err)
		}
		err = initTables(pathToSchema)
		if err != nil {
			return errors.Join(fmt.Errorf("couldn't initialise tables"), err)
		}
	}
	return err
}

func initConnection() error {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))

	// open database
	var err error
	pool, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}

	// check db
	return pool.Ping()
}

func initTables(pathToSchema string) error {
	schema, err := os.ReadFile(pathToSchema)
	if err != nil {
		return err
	}

	_, err = pool.Exec(string(schema))
	if err != nil {
		return err
	}
	return nil
}
