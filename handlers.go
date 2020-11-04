package main

import (
	"net/http"
	"fmt"
	"teste-docker/db"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

func Birds(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, db.Birds())
}