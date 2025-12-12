package shell

import (
	cmd "gsh/internal/command"
	cfg "gsh/internal/config"
	"os"
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

// ShellContext interface implementation
func (s *Shell) GetCurrentDir() string {
	return s.currentDir
}

func (s *Shell) SetCurrentDir(dir string) {
	s.currentDir = dir
}

func (s *Shell) GetEnv(key string) string {
	return s.env[key]
}

func (s *Shell) SetEnv(key, value string) {
	s.env[key] = value
}

func (s *Shell) GetConfig() interface{} {
	return s.config
}

func NewShell(config *cfg.Config) *Shell {
	dir, _ := os.Getwd()
	return &Shell{
		name:       DefaultShellName,
		user:       "user",
		currentDir: dir,
		env:        make(map[string]string),
		config:     config,
	}
}

func New() *Shell {
	var conf *cfg.Config = cfg.NewDeafultConfig()
	return NewShell(conf)
}
