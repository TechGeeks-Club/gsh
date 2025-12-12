package builtins

import (
	"fmt"
	"gsh/internal/command"
	"os"
)

func Pwd(sh ShellContext, c command.Command) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Print(dir)
	return nil
}
