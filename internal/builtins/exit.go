package builtins

import (
	"gsh/internal/command"
	"os"
)

func Exit(sh ShellContext, c command.Command) error {
	os.Exit(0)
	return nil
}
