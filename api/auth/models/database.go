package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func init() {
	envErr := godotenv.Load()

	if envErr != nil {
		fmt.Print(envErr)
		os.Exit(1)
	}

	// @TODO: Drop all db's, create new db for project, and setup db from AutoMigrate fn below.
	// Check tutorial for examples on how to setup AutoMigrate. anthonyscinocco is the db that has the users table
	// currently.
	u := os.Getenv("db_user")
	// p := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", dbHost, u, dbName)

	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	db = conn
	db.Debug().AutoMigrate()
}

func GetDB() *gorm.DB {
	return db
}
