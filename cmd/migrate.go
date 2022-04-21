package cmd

import (
	"gin-blog/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use: "migrate",

	Short: "Migrate tables to database",

	Long: `Migrate tables to database`,

	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.Migrate()
	},
}
