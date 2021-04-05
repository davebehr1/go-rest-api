package cmd

import (
	"context"
	"fmt"
	"log"
	"lxdAssessmentServer/ent/migrate"
	"lxdAssessmentServer/pkg"
	"lxdAssessmentServer/pkg/db"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Creates database tabels in postgres",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := pkg.GetConfig()

		client, _, err := db.NewClient(cfg.Database)

		if err != nil {
			log.Fatalf("failed to create a database connection: %v", err)
		}

		defer client.Close()

		if err := client.Schema.Create(context.Background(), migrate.WithDropColumn(true)); err != nil {
			log.Fatalf("failed to create schema tables in db: %v", err)
		}

		fmt.Println("completed migration")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
