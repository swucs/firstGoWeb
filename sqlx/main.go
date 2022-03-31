package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

const (
	host     string = "localhost"
	port     int    = 5432
	username string = "postgres"
	password string = "1234"
	dbname   string = "postgres"
)

func main() {

	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}

	//ql := &QueryLogger{db, &log.Logger{}}

	//struct에서 'db'대신에 'json' 태그로 매핑한다.
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)

	fmt.Println("DB 연결완료")

	//트랜잭션 시작
	//tx := db.MustBegin()
	//CreateCustomer(tx)

	SelectCustomer(db)
}
