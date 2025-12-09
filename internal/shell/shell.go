package shell

import (
	cmd "gsh/internal/command"
	cfg "gsh/internal/config"
)

type Shell struct {
	currentDir string
	rawInput   string
	command    cmd.Command
	history    []string
	env        map[string]string
	config     *cfg.Config
}

func (s *Shell) Command() cmd.Command {
	return s.command
}

func NewShell(config *cfg.Config) *Shell {
	return &Shell{
		currentDir: "/",
		env:        make(map[string]string),
		config:     config,
	}
}
