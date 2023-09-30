package postgre

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func connectToDB() {

}

func getUser() {

}

type test struct {
	db *sql.DB
}

func (t *test) RegUser() string {
	//TODO implement me
	panic("implement me")
}

func (t *test) GetUser() string {
	//TODO implement me
	panic("implement me")
}

func (t *test) CheckUser() string {
	//TODO implement me
	panic("implement me")
}

// TODO в чем разница тут если вернуть с указателем и без
func NewPgConnect() (*test, error) {
	connStr := "host=localhost port=5444 user=postgres password=12345 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &test{db: db}, nil
}

// TODO в чем разница тут если будет с указателем и без
func (t *test) AddUser() string {

	return ""
}

func AddUser() string {
	return ""
}
