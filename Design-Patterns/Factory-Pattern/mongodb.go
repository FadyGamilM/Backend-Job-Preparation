package main

import (
	"errors"
	"log"
)

// the mongodb concrete impl of our Database abstraction
type mongoDB struct {
	db map[int]string
}

// implement the Database interface via the MongoDB type
func (m mongoDB) GetByID(id int) (string, error) {
	if _, ok := m.db[id]; !ok {
		return "", errors.New("not found")
	}
	log.Println("From Mongodb database engine")
	return m.db[id], nil
}

func (m *mongoDB) Insert(id int, data string) {
	m.db[id] = data
}
