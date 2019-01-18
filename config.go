package main

import (
	"context"

	"github.com/urfave/cli"
)

type ConfigSource interface {
	List(ctx context.Context) ([]string, error)
}

const metadataConfigSourceKey = "configSource"

func setConfigSource(c *cli.Context, src ConfigSource) {
	c.App.Metadata[metadataConfigSourceKey] = src
}

func configSourceFromContext(c *cli.Context) ConfigSource {
	if src, ok := c.App.Metadata[metadataConfigSourceKey].(ConfigSource); ok {
		return src
	}
	return nil
}
