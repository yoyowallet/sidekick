package main

import (
	"os"
	"os/exec"
)

type Process struct {
	*exec.Cmd
}

func NewProcess(name string, arg ...string) *Process {
	proc := new(Process)
	proc.Cmd = exec.Command(name, arg...)

	// Defaults
	proc.Env = os.Environ()
	proc.Stdin = os.Stdin
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr

	return proc
}

func (p *Process) Start() error {
	return p.Cmd.Start()
}
