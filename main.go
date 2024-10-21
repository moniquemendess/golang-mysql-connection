package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// environment variable configuration .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// variable password
	dbPassword := os.Getenv("DB_PASSWORD")
	db, err := sql.Open("mysql", "root:"+dbPassword+"@tcp(127.0.0.1:3306)/librosdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Sucess!")
}
