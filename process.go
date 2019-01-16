package main

import (
	"context"
	"os"
	"os/exec"
)

type Process struct {
	*exec.Cmd
}

func NewProcess(name string, arg ...string) *Process {
	p := new(Process)
	p.Cmd = exec.Command(name, arg...)

	// Defaults
	p.Env = os.Environ()
	p.Stdin = os.Stdin
	p.Stdout = os.Stdout
	p.Stderr = os.Stderr

	return p
}

func (p *Process) AppendEnvFromSource(src ConfigSource) error {
	items, err := src.List(context.Background())
	if err != nil {
		return err
	}
	p.Env = append(p.Env, items...)
	return nil
}

func (p *Process) Start() error {
	return p.Cmd.Start()
}
