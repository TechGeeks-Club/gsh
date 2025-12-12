package builtins

import (
	"fmt"
	cmd "gsh/internal/command"
)

func Clear(sh ShellContext, c cmd.Command) error {
	fmt.Print("\033[2J\033[H")
	return nil
}
