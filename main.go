package main

import (
	"github.com/dwburke/mockapi/cmd"

	"github.com/dwburke/mockapi/api"
	"github.com/dwburke/mockapi/cron"
	"github.com/dwburke/mockapi/logging"
)

func main() {
	defer cron.Shutdown()
	defer logging.Cleanup()

	cmd.Execute()
	api.Shutdown()
}
