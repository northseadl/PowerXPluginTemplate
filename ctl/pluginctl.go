package main

import (
	"PowerXPluginTemplate/ctl/plugin"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "pluginctl",
		Usage: "pluginctl",
		Commands: []*cli.Command{
			plugin.BuildCmd,
			plugin.NewCmd,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
