package main

import (
	"log"
	"reflect"
)

func main() {
	dbEngine1 := "mongodb"
	dbEngine2 := "postgres"

	mongodbInstance := DatabaseFactory(dbEngine1)
	postgresInstance := DatabaseFactory(dbEngine2)

	mongodbInstance.Insert(1, "fady")
	mongodbInstance.Insert(2, "mina")
	result, _ := mongodbInstance.GetByID(1)
	// log.Println(reflect.TypeOf(mongodbInstance).Name())
	log.Println(reflect.TypeOf(&mongodbInstance).Elem().Name()) // this will return the interface that this type implements
	log.Println(result)

	postgresInstance.Insert(1, "marwan")
	postgresInstance.Insert(2, "sa3ed")
	result, _ = postgresInstance.GetByID(2)
	// log.Println(reflect.TypeOf(postgresInstance).Name())
	log.Println(reflect.TypeOf(&postgresInstance).Elem().Name()) // this will return the interface that this type implements
	log.Println(result)
}

// the interface for our database repo
type Database interface {
	GetByID(id int) (string, error)
	Insert(id int, data string)
}
