package main

import (
	"github.com/jinzhu/gorm"
)

func getDBConnection(dsn string) *gorm.DB {
	// TODO: create database connection from the dsn
	gorm.Open("postgres", dsn)
	return nil
}

func main() {
	// TODO:
	//  create database connection
	//  define connection parameters
	//  create transaction from the database object
	//  use transaction to create a new record in the database
	//  commit the transaction

}
