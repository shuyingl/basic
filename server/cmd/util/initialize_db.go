package main

// This is just a simple script to create the database schema and insert a test user
// Don't use this in production, it's just for debugging purposes

import (
	"database/sql"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"server/embeds"
	"sort"

	_ "github.com/lib/pq" // Importing the driver for PostgreSQL
)

func main() {
	dbNamePtr := flag.String("db_name", "test", "name of the database to create")
	flag.Parse()

	err := SetupSchema(*dbNamePtr)
	if err != nil {
		log.Printf("Error setting up schema: %s", err)
		return
	}
}

func SetupSchema(dbName string) error {
	connStr := "user=postgres dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	var exists bool
	db.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT datname FROM pg_catalog.pg_database WHERE datname = '%s')", dbName)).
		Scan(&exists)
	if !exists {
		_, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			return err
		}
	}

	connStr = fmt.Sprintf("user=postgres dbname=%s sslmode=disable", dbName)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// Create the database schema
	entries, _ := fs.ReadDir(embeds.DatabaseSchema, "sqls")
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})
	for _, entry := range entries {
		fileContent, _ := fs.ReadFile(embeds.DatabaseSchema, "sqls/"+entry.Name())
		_, err := db.Exec(string(fileContent))
		if err != nil {
			return err
		}
	}
	return nil
}
