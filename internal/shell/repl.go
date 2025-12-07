package shell

import (
	cmd "/gsh/internal/command"
	"bufio"
	"fmt"
	"os"
)

func readInput(shell *Shell) (cmd.Command, error) {
	var err error
	var command string

	command, err = bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return cmd.Command{}, fmt.Errorf("error reading input: %w", err)
	}

	shell.rawInput = command
	shell.commands = cmd.Parse(command)
	shell.history = append(shell.history, command)

	return shell.commands, nil
}
