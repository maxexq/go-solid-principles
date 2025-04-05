package main

import "fmt"

type (
	PostgresDb struct{}
	MockDb     struct{}
)

func InsertPlayerItemPostgres(db *PostgresDb) {
	db.insert()
}

func (db *PostgresDb) insert() {
	fmt.Println("Inserting into PostgresDb")
}

func InsertPlayerItemMock(db *MockDb) {
	db.insert()
}

func (db *MockDb) insert() {
	fmt.Println("Inserting into MockDb")
}

func main() {
	postgresDb := &PostgresDb{}
	InsertPlayerItemPostgres(postgresDb)

	mockDb := &MockDb{}
	InsertPlayerItemMock(mockDb)
}

// dependency inversion principle
// In this example, we have two database types: `PostgresDb` and `MockDb`.
// The `InsertPlayerItemPostgres` and `InsertPlayerItemMock` functions are responsible for inserting player items into their respective databases.
// This violates the Dependency Inversion Principle because the high-level modules (the functions) depend on low-level modules (the database types) directly.
// Instead, we should use interfaces to abstract the database operations and allow for easier testing and flexibility in the future.
