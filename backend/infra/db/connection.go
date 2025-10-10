package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string{
	return "user=postgres password=1212 host=localhost port=5432 dbname=ecommerce sslmode=disable"
}


func NewConnection() (*sqlx.DB, error){
	dbSource := GetConnectionString()

	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil{
		fmt.Println(err)
		return nil, err
	}
	return dbCon, nil
}