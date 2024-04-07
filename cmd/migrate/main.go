package main

import (
	"context"
	"log"

	"example.com/application/config"
	"example.com/application/pkg/datastore"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

var MigrateCreateCmd = &cobra.Command{
	Use:       "create",
	Short:     "Migration tool",
	ValidArgs: []string{"name"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalln("Must provide the 'name' of the migration as first argument")
		}
		name := args[0]

		cfg := config.GetConfig()
		db, err := datastore.CreateConnection(context.Background(), cfg.DbUri)
		if err != nil {
			log.Fatalln("Error setting up connection", err)
		}
		defer db.Close()

		if err := goose.Create(db, "schemas/migrations", name, "sql"); err != nil {
			log.Fatalln("Error creating migrations", err)
		}

	},
}

func main() {
	config.ReadConfig()

	MigrateCreateCmd.Execute()
}
