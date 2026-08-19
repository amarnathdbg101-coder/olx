package main

import (
	"fmt"
	"log"
	"os"

	"github.com/amarnathdbg101-coder/olx/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: migrate <up | down>")
		return
	}
	m, err := migrate.New("file://internal/migrations",config.MustLoad().DatabaseUrl)
	if err != nil {
		log.Fatalf("migration error:%v", err)
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			if err == migrate.ErrNoChange {
				fmt.Println("No new migrations to apply")
			} else {
				log.Fatalf("migration up failed: %v", err)
			}
		} else {
			fmt.Println("Successfully applied migrations")
		}
	case "down":
		if err := m.Steps(-1); err != nil {
			if err == migrate.ErrNoChange {
				fmt.Println("No migrations to roll back")
			} else {
				log.Fatalf("migration down failed: %v", err)
			}
		} else {
			fmt.Println("Successfully rolled back migration")
		}
	default:
		log.Fatalf("unknown command :%s", os.Args[1])
	}
}
