package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

const metadataConfigSource = "configSource"

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config-source",
			EnvVar: "SIDEKICK_CONFIG_SOURCE",
			Value:  "dynamodb",
		},
		cli.StringFlag{
			Name:   "config-table",
			EnvVar: "SIDEKICK_CONFIG_TABLE",
		},
	}
	app.Before = func(c *cli.Context) error {
		var configSource ConfigSource
		switch c.String("config-source") {
		case "dynamodb":
			configSource = &DynamoDBConfigSource{
				Table: c.String("config-table"),
				Key:   "common", // TODO: parameterise
			}
		default:
			return cli.NewExitError("couldn't find that config source type", 2)
		}

		c.App.Metadata[metadataConfigSource] = configSource

		return nil
	}
	app.Commands = []cli.Command{
		commandStart,
		commandEnv,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
