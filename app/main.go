package main

import (
	"database/sql"
	"github.com/paulmarie/univesity/grade_app/api"
	grade_database "github.com/paulmarie/univesity/grade_app/persistence/sqlc"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver     = "postgres"
	dbSources    = "postgresql://user:password@university.grade.database.fr:5432/grade_dev?sslmode=disable"
	serverAdress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSources)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := grade_database.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
