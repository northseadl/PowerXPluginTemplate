package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "pluginctl",
		Usage: "pluginctl",
		Commands: []*cli.Command{
			{
				Name: "build",
				Action: func(c *cli.Context) error {
					builder, err := NewBuilder(".")
					if err != nil {
						return err
					}
					if err = builder.CheckDependencies(); err != nil {
						return err
					}
					if err = builder.BuildBackend("backend"); err != nil {
						return err
					}
					if err = builder.BuildFrontend("frontend"); err != nil {
						return err
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
