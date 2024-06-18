package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "/Users/makumar/Documents/workroom/gows/discovery/nios.db")
	if err != nil {
		fmt.Printf("Error opening sqlite db %v", err)
		panic(err)
	}
	defer db.Close()

	// List all table names in the database
	tables := listTables(db)

	// Delete all tables
	deleteTables(db, tables)

	// Read the content of the SQL file
	sqlFile := "/Users/makumar/Downloads/nios_schema.sql"
	sqlContent, err := os.ReadFile(sqlFile)
	if err != nil {
		fmt.Errorf("Error reading nios schema file %v", err)
		panic(err)
	}

	// Split the content into separate SQL statements
	sqlStatements := string(sqlContent)
	statements := splitStatements(sqlStatements)

	// Execute each SQL statement
	for _, stmt := range statements {
		_, err := db.Exec(stmt)
		if err != nil {
			fmt.Errorf("Error executing statement: %v", err)
			panic(err)
		}
	}
	return
}

func splitStatements(sqlContent string) []string {
	// Split the content into individual SQL statements
	// Assumes statements are separated by semicolon followed by newline
	statements := []string{}
	current := ""
	for _, char := range sqlContent {
		current += string(char)
		if char == ';' {
			statements = append(statements, current)
			current = ""
		}
	}
	return statements
}

func listTables(db *sql.DB) []string {
	tables := []string{}
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		fmt.Errorf("%v", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			fmt.Errorf("%v", err)
		}
		tables = append(tables, tableName)
	}

	return tables
}

func deleteTables(db *sql.DB, tableNames []string) {
	for _, tableName := range tableNames {
		_, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName))
		if err != nil {
			fmt.Errorf("Error deleting table %s: %v", tableName, err)
			return
		}
	}
	fmt.Println("Table deletion successfully.")
}
