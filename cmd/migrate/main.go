package main

import (
	"fmt"
	"log"
	"os"

	"example.com/application/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migration tool",
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $XDG_CONFIG_HOME/application.toml)")
	cobra.OnInitialize(func() {
		cfgFile := rootCmd.Flags().Lookup("config")
		if cfgFile.Value.String() != "" {
			log.Println("Config file set", cfgFile.Value.String())

			if err := config.ReadFile(cfgFile.Value.String()); err != nil {
				log.Fatalln("Error reading config file", err)
			}
		} else {
			if err := config.FindAndReadFile(); err != nil {
				log.Fatalln("Error reading config file", err)
			}
		}
	})
	rootCmd.AddCommand(migrateCreateCmd)
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
