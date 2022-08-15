package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var validate = validator.New()

func main() {
	log.Println("starting cars backend")

	dbConfig := getDBConfig()

	err := migrateDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		dbConfig.Host, dbConfig.User, dbConfig.Pass, dbConfig.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(db)
	}

	r := createRouter(db)

	log.Println("web server is ready")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

func createRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/cars", func(r chi.Router) {
		r.Get("/", listCarsHandler(db))
		r.Get("/{id}", getCarHandler(db))
		r.Post("/{id}", createCarHandler(db))
	})

	return r
}
