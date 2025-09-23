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
	dsn := os.Getenv("DB_DSN") // e.g. postgres://user:pass@localhost:5432/schema_db?sslmode=disable
	db, err := sql.Open("postgres", dsn)
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

	log.Println("Schema Versioner API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
