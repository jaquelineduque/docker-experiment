package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}


func main() {
	router := NewRouter()
	port:= os.Getenv("APP_PORT")
	if (port == ""){
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}