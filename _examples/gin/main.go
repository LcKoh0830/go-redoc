package main

import (
	"github.com/LcKoh0830/go-redoc"
	ginredoc "github.com/LcKoh0830/go-redoc/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./openapi.json",
		SpecPath:    "/openapi.json",
		DocsPath:    "/docs",
	}

	r := gin.New()
	r.Use(ginredoc.New(doc))

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(r.Run(":8000"))
}
