package shell

import (
	"bufio"
	"fmt"
	blt "gsh/internal/builtins"
	cmd "gsh/internal/command"
	"os"
)

func (shell *Shell) readInput() (cmd.Command, error) {
	var err error
	var cmmnd string

	cmmnd, err = bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return cmd.Command{}, fmt.Errorf("error reading input: %w", err)
	}
	shell.rawInput = cmmnd
	shell.command = cmd.Parse(cmmnd)
	shell.history = append(shell.history, cmmnd)

	return shell.command, nil
}

func (shell *Shell) eval() { // to fix  eval
	var cmnd cmd.Command = shell.Command()
	fnc, exist := blt.Builtins[cmnd.Base()]
	if exist {
		fnc(cmnd)
	}
	/// ...tobe continued
	// return nil
}

func (shell *Shell) errorHandler(err error) {
	if err != nil {
		fmt.Fprintf(os.Stdin, "Error")
	}

}

func (s *Shell) REPL() {
}
