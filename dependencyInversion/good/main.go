package main

import "fmt"

type (
	Database interface {
		Insert()
	}

	PostgresDb struct{}
	MockDb     struct{}
)

func NewPostgresDb() Database {
	return &PostgresDb{}
}

func NewMockDb() Database {
	return &MockDb{}
}

func (db *PostgresDb) Insert() {
	fmt.Println("Inserting into PostgresDb")
}

func (db *MockDb) Insert() {
	fmt.Println("Inserting into MockDb")
}

// InsertPlayerItem is a high-level function that depends on the Database interface
// instead of a specific database implementation. This allows for flexibility and easier testing.
func InsertPlayerItem(db Database) {
	db.Insert()
}

func main() {
	postgresDb := NewPostgresDb()
	InsertPlayerItem(postgresDb)

	mockDb := NewMockDb()
	InsertPlayerItem(mockDb)
}

// dependency inversion principle
// In this example, we have two database types: `PostgresDb` and `MockDb`. The `InsertPlayerItem` function is responsible for inserting player items into the database.
// By using the `Database` interface, we can easily switch between different database implementations without modifying the function itself.
// This adheres to the Dependency Inversion Principle, allowing for better flexibility and testability in our code.
