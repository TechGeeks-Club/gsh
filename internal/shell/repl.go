package shell

import (
	"bufio"
	"fmt"
	blt "gsh/internal/builtins"
	cmd "gsh/internal/command"
	"os"
)

func readInput(shell *Shell) (cmd.Command, error) {
	var err error
	var cmmnd string

	cmmnd, err = bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return cmd.Command{}, fmt.Errorf("error reading input: %w", err)
	}
	shell.rawInput = cmmnd
	shell.commands = cmd.Parse(cmmnd)
	shell.history = append(shell.history, cmmnd)

	return shell.commands, nil
}

func eval(shell *Shell, command cmd.Command) error {
	fnc, exist := blt.Builtins[command.Base()]
	if exist {
		fnc(command)
	}
	/// ...tobe continued
	return nil
}
