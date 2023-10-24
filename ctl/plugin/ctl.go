package plugin

import "github.com/urfave/cli/v2"

var BuildCmd = &cli.Command{
	Name:  "build",
	Usage: "build plugin or backend API",
	Subcommands: []*cli.Command{
		{
			Name:  "plugin",
			Usage: "build plugin",
			Action: func(c *cli.Context) error {
				builder, err := NewBuilder(".")
				if err != nil {
					return err
				}
				if err = builder.Check(); err != nil {
					return err
				}
				if err = builder.BuildBackend("backend"); err != nil {
					return err
				}
				if err = builder.BuildFrontend("frontend"); err != nil {
					return err
				}
				if err = builder.BuildExtra(); err != nil {
					return err
				}
				if err = builder.Package(); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "backend-api",
			Usage: "build backend API",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "input",
					Aliases: []string{"i"},
					Value:   "backend/plugin.api",
					Usage:   "input file for backend API",
				},
				&cli.StringFlag{
					Name:    "output",
					Aliases: []string{"o"},
					Value:   "backend",
					Usage:   "output directory for backend API",
				},
			},
			Action: func(c *cli.Context) error {
				input := c.String("input")
				output := c.String("output")
				return buildAPI(input, output)
			},
		},
	},
}

var NewCmd = &cli.Command{
	Name:  "new",
	Usage: "new plugin",
	Action: func(c *cli.Context) error {
		starter := NewStarter()
		if err := starter.Start(); err != nil {
			return err
		}
		return nil
	},
}
