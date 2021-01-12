package main

import (
	"api-test/db"
	"api-test/routes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println("Server running in port " + port)
	db.DB = db.SetupDB()
	defer db.DB.Close()
	log.Fatal(http.ListenAndServe(":"+port, routes.RouterConfig()))
}
