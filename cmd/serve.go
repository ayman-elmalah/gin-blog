package cmd

import (
	"gin-blog/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use: "serve",

	Short: "Serve application",

	Long: `Application will be served on port that written in .env.`,

	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.Serve()
	},
}
