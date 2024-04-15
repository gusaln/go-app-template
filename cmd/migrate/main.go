package main

import (
	"fmt"
	"os"

	"example.com/application/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migration tool",
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $XDG_CONFIG_HOME/keepingtabs.toml)")
	cobra.OnInitialize(func() {
		cfgFile := rootCmd.Flags().Lookup("config")
		if cfgFile.Value.String() != "" {
			viper.SetConfigFile(cfgFile.Value.String())
		}

		config.ReadConfigFile()
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
