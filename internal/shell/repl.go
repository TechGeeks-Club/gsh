package shell

import (
	"bufio"
	"fmt"
	"os"

	blt "gsh/internal/builtins"
	cmd "gsh/internal/command"
	cfg "gsh/internal/config"
)

const (
	COMMAND_NOT_FOUND int8 = 0
)

func (shell *Shell) promptBuilder() {
	if shell.config.ColorMode {
		fmt.Fprintf(shell.config.Stdout(), " %s\n", shell.currentDir)
		fmt.Fprintf(shell.config.Stdout(), "%sgsh>%s ", cfg.Cyan, cfg.Reset)
	} else {
		fmt.Fprintf(shell.config.Stdout(), " %s\n", shell.currentDir)
		fmt.Fprintf(shell.config.Stdout(), "gsh> ")
	}
}

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

func (shell *Shell) eval() int8 {
	var cmnd cmd.Command = shell.Command()
	fnc, exist := blt.Builtins[cmnd.Base()]
	if exist {
		fnc(cmnd)
	} else {
		shell.errorHandler(COMMAND_NOT_FOUND)
		return 1
	}
	return 0
}

func (shell *Shell) errorHandler(err int8) {
	if err == 0 {
		fmt.Fprintf(shell.config.Stdout(), "%s: %s: command not found", shell.name, shell.command.Raw())
	} else {
		fmt.Fprintf(shell.config.Stdout(), "%s: %s: command not found", shell.name, shell.command.Raw())
	}

}

func (shell *Shell) repl() {
	var code int8
	for {
		shell.promptBuilder()
		shell.readInput()
		code = shell.eval()
		if code == 1 {
			break
		}
	}
}

func Run() {
	var shell *Shell = New()
	shell.repl()
}
