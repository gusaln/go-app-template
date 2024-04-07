package migrate

import (
	"context"
	"embed"
	"log"

	"example.com/application/config"
	"example.com/application/pkg/datastore"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

var MigrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Migration tool",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()

		db, err := datastore.CreateConnection(context.Background(), cfg.DbUri)
		if err != nil {
			log.Fatalln("Error setting up connection", err)
		}

		goose.SetBaseFS(embedMigrations)
		if err := goose.Up(db, "migrations"); err != nil {
			log.Fatalln("Error migrating", err)
		}

		log.Println("Migrations up to date")
	},
}

var MigrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Migration tool",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()

		db, err := datastore.CreateConnection(context.Background(), cfg.DbUri)
		if err != nil {
			log.Fatalln("Error setting up connection", err)
		}

		goose.SetBaseFS(embedMigrations)
		if err := goose.Down(db, "migrations"); err != nil {
			log.Fatalln("Error migrating", err)
		}

		log.Println("Migrations rolled back")
	},
}
