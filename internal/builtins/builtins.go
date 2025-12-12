package builtins

import "gsh/internal/command"

type ShellContext interface {
	GetCurrentDir() string
	SetCurrentDir(dir string)
	GetEnv(key string) string
	SetEnv(key, value string)
	GetConfig() interface{}
}

var Builtins = map[string]func(sh ShellContext, c command.Command) error{
	"echo":  Echo,
	"help":  Help,
	"exit":  Exit,
	"pwd":   Pwd,
	"clear": Clear,
	"cd":    Cd,
	"ls":    Ls,
}
