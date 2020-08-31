package main

import (
	"os"
	"testing"
	"log"
)

var a App

func TextMain(m *testing.M){
	a.Initialize(os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

}

func ensureTableExists(){
	_, err:= a.DB.Exec(tableCreationQuery)
	if (err != nil){
		log.Fatal(err)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
id SERIAL,
name TEXT NOT NULL,
price NUMERIC(10,2) not null default 0.00,
CONSTRAINT products_pkey PRIMARY KEY (id)
)`

func clearTable(){
	a.DB.Exec(`DELETE FROM products`)
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}