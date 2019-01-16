package main

import (
	"github.com/urfave/cli"
)

var commandEnv = cli.Command{
	Name: "env",
	Action: func(c *cli.Context) error {
		var err error

		proc := NewProcess("env")
		proc.AppendEnvSource(c.App.Metadata[metadataConfigSource].(ConfigSource))

		err = proc.Start()
		if err != nil {
			return err
		}

		err = proc.Wait()
		if err != nil {
			return err
		}

		return nil
	},
}
