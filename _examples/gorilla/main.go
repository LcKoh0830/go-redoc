package main

import (
	"net/http"

	"github.com/LcKoh0830/go-redoc"
	"github.com/gorilla/mux"
)

func main() {
	doc := &redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./openapi.json",
		SpecPath:    "/docs/openapi.json",
	}

	r := mux.NewRouter()
	r.PathPrefix("/docs").Handler(doc.Handler())

	println("Documentation served at http://127.0.0.1:8000/docs")
	http.ListenAndServe(":8000", r)
}
