package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://cmd/migrate/migrations",
		"postgres://jet:2004@localhost:5432/nexora_DB?sslmode=disable",
	)
	if err != nil {
		log.Fatal("error in initialize migration:", err)
	}

	cmd := os.Args[len(os.Args)-1]

	if cmd == "up" {
		if err := m.Up(); err != nil {
			if err == migrate.ErrNoChange {
				log.Println("no new migrations to apply")
			} else {
				log.Fatal("migrate up error:", err)
			}
		} else {
			log.Println("migrate up success")
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil {
			if err == migrate.ErrNoChange {
				log.Println("no migrations to rollback")
			} else {
				log.Fatal("migrate down error:", err)
			}
		} else {
			log.Println("migrate down success")
		}
	}
}
