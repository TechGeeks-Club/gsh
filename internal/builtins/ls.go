package builtins

import (
	"fmt"
	cmd "gsh/internal/command"
	cfg "gsh/internal/config"
	"os"
)

func Ls(sh ShellContext, c cmd.Command) error {
	var path string = "."
	if len(c.Args()) > 0 && c.Args()[0] != "" {
		path = c.Args()[0]
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.Name()[0] == '.' {
			continue
		}
		if entry.IsDir() {
			fmt.Printf("%s%s%s\n", cfg.Blue, entry.Name(), cfg.Reset)
		} else {
			fmt.Printf("%s\n", entry.Name())

		}

	}
	return nil
}
