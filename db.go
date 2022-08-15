package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DBConfig struct {
	Host string
	Name string
	User string
	Pass string
}

func getDBConfig() *DBConfig {
	log.Println("fetching db configurations via environment")

	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		log.Fatal("DB_HOST env variable is not set")
	}

	name, ok := os.LookupEnv("DB_NAME")
	if !ok {
		log.Fatal("DB_NAME env variable is not set")
	}

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		log.Fatal("DB_USER env variable is not set")
	}

	pass, ok := os.LookupEnv("DB_PASS")
	if !ok {
		log.Fatal("DB_PASS env variable is not set")
	}

	return &DBConfig{
		Host: host,
		Name: name,
		User: user,
		Pass: pass,
	}
}

func migrateDB(config *DBConfig) error {
	log.Println("migrating the database")

	conn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		config.User, config.Pass, config.Host, config.Name)

	m, err := migrate.New("file://db/migrations", conn)
	if err != nil {
		return err
	}

	// Migrate up igoring if already migrated
	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	return err
}
