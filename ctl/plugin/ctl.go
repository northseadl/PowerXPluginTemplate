package plugin

import "github.com/urfave/cli/v2"

var BuildCmd = &cli.Command{
	Name:  "build",
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
}
