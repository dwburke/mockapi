package cmd

import (
	"github.com/spf13/cobra"

	"github.com/dwburke/mockapi/api"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the mock api",
	Long:  `Start the mock api`,
	Run: func(cmd *cobra.Command, args []string) {
		api.Run()
		<-api.ShutdownCh
	},
}
