package main

// the factory method for returning any type of database engine that implements the abstract Database type
func DatabaseFactory(dbEngine string) Database {
	switch dbEngine {
	case "mongodb":
		return &mongoDB{db: make(map[int]string)}
	case "postgres":
		return &postgreSQL{db: make(map[int]string)}
	default:
		return nil
	}
}
