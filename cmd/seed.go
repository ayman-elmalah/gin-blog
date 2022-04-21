package cmd

import (
	"gin-blog/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use: "seed",

	Short: "Seed data to database",

	Long: `Seed data to database`,

	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.Seed()
	},
}
