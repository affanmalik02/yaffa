package main

import (
	"log"
)

func GenerateOpenAPISpec() error {
	// TODO: generate OpenAPI 3.0 spec from handlers/types
	log.Println("Generating OpenAPI spec...")
	return nil
}

func ServeSwaggerFiles(dir string) error {
	// TODO: serve static swagger UI files
	log.Printf("Serving Swagger UI from %s\n", dir)
	return nil
}

func main() {
	if err := GenerateOpenAPISpec(); err != nil {
		log.Fatalf("Failed to generate OpenAPI spec: %v", err)
	}
}