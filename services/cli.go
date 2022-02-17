package services

import "os/exec"

type Cli struct{}

func (c Cli) Run(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).Output()
}
