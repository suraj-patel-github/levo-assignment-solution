package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"levo-schema-service/internal/storage"
	"levo-schema-service/pkg/repository"
	"levo-schema-service/pkg/service"
	"levo-schema-service/pkg/transport"

	_ "github.com/lib/pq"
)

func main() {
	dbString := os.Getenv("POSTGRES_CONNECTION_STRING")
	if dbString == "" {
		log.Fatal("DB connection is not provided")
	}
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.New(db)
	fs := storage.NewFileStore("./uploads")
	svc := service.New(repo, fs)

	mux := http.NewServeMux()
	mux.Handle("/upload", transport.UploadHandler(svc))
	mux.Handle("/latest", transport.LatestHandler(svc))
	mux.Handle("/version", transport.VersionHandler(svc))

	log.Println("Schema Specs Validation API on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
