package cmd

import (
	"fmt"
	"os"

	"github.com/dwburke/vipertools"
	"github.com/spf13/cobra"

	"github.com/dwburke/mockapi/cron"
	"github.com/dwburke/mockapi/logging"
)

var configDir string

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(logging.InitLogging)
	cobra.OnInitialize(cron.Run)
	rootCmd.PersistentFlags().StringVar(&configDir, "config-dir", "etc/", "directory to read config files from (default is ./etc/)")
}

var rootCmd = &cobra.Command{
	Use:   "dburke-things",
	Short: "dburke-things is a thing",
	Long:  `Love me`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if err := vipertools.ReadDir(configDir); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
