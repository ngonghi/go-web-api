package main

import (
	"fmt"
	"github.com/ngonghi/vian-backend/cmd/app"
	"github.com/ngonghi/vian-backend/cmd/db"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "CLI for vian",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "App server related",
}

var appServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run application server",
	RunE:  app.Serve,
}

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Manipulate the database",
}

var dbMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run the migrations",
	RunE:  db.Migrate,
}

var dbSeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run the migrations",
	RunE:  db.Seed,
}

var dbCreateMigrationCmd = &cobra.Command{
	Use:   "create_migration",
	Short: "Run the create migration file",
	RunE:  db.CreateMigration,
}

func init() {
	rootCmd.AddCommand(appCmd)
	appCmd.AddCommand(appServeCmd)
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(dbMigrateCmd)
	dbCmd.AddCommand(dbSeedCmd)
	dbCmd.AddCommand(dbCreateMigrationCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
