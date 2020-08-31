/*
Following tutorial from
https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
 */
package main
import (

	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lib/pq" //to work with postgres
	"log"
)

type App struct {
	Router *mux.Router
	DB *sql.DB
}

func (a *App) Initialize(username, password, dbname string){
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",username,password,dbname)
	var err error
	a.DB, err = sql.Open("postgres",connectionString)
	if err != nil{
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
}

func (a *App) Run(address string){}

