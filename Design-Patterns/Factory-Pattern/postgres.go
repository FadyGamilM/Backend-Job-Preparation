package main

import (
	"errors"
	"log"
)

// the postgres concrete impl for our Database abstraction
type postgreSQL struct {
	db map[int]string
}

// implement the Database interface via the PostgreSQL type
func (p postgreSQL) GetByID(id int) (string, error) {
	if _, ok := p.db[id]; !ok {
		return "", errors.New("not found")
	}
	log.Println("From PostgreSQL database engine")
	return p.db[id], nil
}

func (p *postgreSQL) Insert(id int, data string) {
	p.db[id] = data
}
