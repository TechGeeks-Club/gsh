package shell

import (
	"gsh/internal/command"
	cfg "gsh/internal/config"
)

type Shell struct {
	currentDir string
	rawInput   string
	commands   command.Command
	history    []string
	env        map[string]string
	config     *cfg.Config
}

func (s *Shell) REPL() {
}

func NewShell(config *cfg.Config) *Shell {
	return &Shell{
		currentDir: "/",
		env:        make(map[string]string),
		config:     config,
	}
}
