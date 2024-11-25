package cmd

import (
	"github.com/phn00dev/go-crud-temp/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Application run server",
	Long:  `Application run server with config.yml file env`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.Serve()
	},
}
