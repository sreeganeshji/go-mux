package main

import (
	"database/sql"
	"errors"
)


type Product struct {
	name string `json:"name"`
	ID int `json:"id"`
	price float64 `json:"price"`
}

func (p *Product) getPrice(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Product) deleteProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Product) updateProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Product) createProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getProduct(db *sql.DB, start, count int) ([]Product,error ){
	return _, errors.New("Not implemented")
}