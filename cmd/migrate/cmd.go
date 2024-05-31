package main

import (
	"context"
	"log"

	"example.com/application/config"
	"example.com/application/internal/datastore"
	migrate "example.com/application/schemas"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var migrateCreateCmd = &cobra.Command{
	Use:   "create [NAME]",
	Short: "Create a migration",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		cfg := config.Get()
		db, err := datastore.CreateConnection(context.Background(), cfg.DbUri)
		if err != nil {
			log.Fatalln("Error setting up connection", err)
		}
		defer db.Close()

		if err := goose.SetDialect(string(goose.DialectSQLite3)); err != nil {
			log.Fatalln("Error setting dialect:", err)
		}

		if err := goose.Create(db, "schemas/migrations", name, "sql"); err != nil {
			log.Fatalln("Error creating migrations", err)
		}

	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Migrate up",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Get()

		db, err := datastore.CreateConnection(context.Background(), cfg.DbUri)
		if err != nil {
			log.Fatalln("Error setting up connection", err)
		}

		if err := goose.SetDialect(string(goose.DialectSQLite3)); err != nil {
			log.Fatalln("Error setting dialect:", err)
		}

		goose.SetBaseFS(migrate.EmbedMigrations)
		if err := goose.Up(db, "migrations"); err != nil {
			log.Fatalln("Error migrating:", err)
		}

		log.Println("Migrations up to date")
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Migrate down",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Get()

		db, err := datastore.CreateConnection(context.Background(), cfg.DbUri)
		if err != nil {
			log.Fatalln("Error setting up connection", err)
		}

		if err := goose.SetDialect(string(goose.DialectSQLite3)); err != nil {
			log.Fatalln("Error setting dialect:", err)
		}

		goose.SetBaseFS(migrate.EmbedMigrations)
		if err := goose.Down(db, "migrations"); err != nil {
			log.Fatalln("Error migrating:", err)
		}

		log.Println("Migrations rolled back")
	},
}
