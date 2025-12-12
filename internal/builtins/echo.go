package builtins

import (
	"fmt"
	cmd "gsh/internal/command"
)

func Echo(sh ShellContext, c cmd.Command) error {
	fmt.Print(c.Raw()[5:])
	return nil
}
