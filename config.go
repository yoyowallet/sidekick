package main

import "context"

type ConfigSource interface {
	List(ctx context.Context) ([]string, error)
}
