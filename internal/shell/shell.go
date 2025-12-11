package shell

import (
	cmd "gsh/internal/command"
	cfg "gsh/internal/config"
)

type Shell struct {
	name       string
	user       string
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
		name:       DefaultShellName,
		user:       "user",
		currentDir: DefaultCurrentDir,
		env:        make(map[string]string),
		config:     config,
	}
}

func New() *Shell {
	var conf *cfg.Config = cfg.NewDeafultConfig()
	return NewShell(conf)
}
