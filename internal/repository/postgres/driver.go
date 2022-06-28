package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// NewConnection creates a connection to Postgres database.
// Do not forget to call 'db.Close()' in the scope of the function in which it is executed
func NewConnection() *sql.DB {
	db, err := sql.Open("pgx", "host=localhost port=5432 dbname=tasks user=postgres password=Ahojako0")
	if err != nil {
		log.Fatalf("Unable to connect to PostgreSQL database: %s", err)
	}
	//defer db.Close()

	testConn(db)
	return db
}

func testConn(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal("Error testing the database ping.")
	}

	var connection string
	err := db.QueryRow("select 'Successfully connected to the database.'").Scan(&connection)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	fmt.Println(connection)
}
