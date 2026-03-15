package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("warning: failed to load .env file: %v", err)
	}

	databaseURL := os.Getenv("DATABASE_MIGRATION_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_MIGRATION_URL is required")
	}

	migrationsPath, err := filepath.Abs("migrations")
	if err != nil {
		log.Fatalf("failed to resolve migrations path: %v", err)
	}

	migrationsURL := fmt.Sprintf("file://%s", filepath.ToSlash(migrationsPath))

	m, err := migrate.New(migrationsURL, databaseURL)
	if err != nil {
		log.Fatalf("failed to initialize migrator: %v", err)
	}

	defer func() {
		sourceErr, dbErr := m.Close()
		if sourceErr != nil {
			log.Printf("warning: failed to close migration source: %v", sourceErr)
		}
		if dbErr != nil {
			log.Printf("warning: failed to close migration database connection: %v", dbErr)
		}
	}()

	command := "up"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "steps":
		if len(os.Args) < 3 {
			log.Fatal("usage: go run ./cmd/migrator steps <n>")
		}
		n, parseErr := strconv.Atoi(os.Args[2])
		if parseErr != nil {
			log.Fatalf("invalid steps count %q: %v", os.Args[2], parseErr)
		}
		err = m.Steps(n)
	case "force":
		if len(os.Args) < 3 {
			log.Fatal("usage: go run ./cmd/migrator force <version>")
		}
		version, parseErr := strconv.Atoi(os.Args[2])
		if parseErr != nil {
			log.Fatalf("invalid version %q: %v", os.Args[2], parseErr)
		}
		err = m.Force(version)
	case "version":
		version, dirty, versionErr := m.Version()
		if errors.Is(versionErr, migrate.ErrNilVersion) {
			log.Println("no migrations have been applied yet")
			return
		}
		if versionErr != nil {
			log.Fatalf("failed to get migration version: %v", versionErr)
		}

		log.Printf("current migration version: %d (dirty=%t)", version, dirty)
		return
	default:
		log.Fatalf("unknown command %q. valid commands: up, down, steps, force, version", command)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Println("no migration changes")
		return
	}

	if err != nil {
		log.Fatalf("migration command %q failed: %v", command, err)
	}

	log.Printf("migration command %q completed successfully", command)

}
