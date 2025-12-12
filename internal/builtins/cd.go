package builtins

import (
	"fmt"
	cmd "gsh/internal/command"
	"os"
)

func Cd(sh ShellContext, c cmd.Command) error {
	if len(c.Args()) > 1 {
		return fmt.Errorf("cd: to many args")

	}

	if c.Args()[0] == "" {
		return fmt.Errorf("cd: missing operand")
	}

	if err := os.Chdir(c.Args()[0]); err != nil {
		return err
	}
	dir, _ := os.Getwd()
	sh.SetCurrentDir(dir)

	return nil
}
