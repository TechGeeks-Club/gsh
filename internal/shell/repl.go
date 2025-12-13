package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	blt "gsh/internal/builtins"
	cmd "gsh/internal/command"
	cfg "gsh/internal/config"
)

const (
	COMMAND_NOT_FOUND int8 = 1
)

func (shell *Shell) promptBuilder() {
	if shell.config.ColorMode {
		fmt.Fprintf(shell.config.Stdout(), "\n%s\n", shell.currentDir)
		fmt.Fprintf(shell.config.Stdout(), "%sgsh>%s ", cfg.Cyan, cfg.Reset)
	} else {
		fmt.Fprintf(shell.config.Stdout(), "\n%s\n", shell.currentDir)
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
	if cmmnd == "" {
		shell.rawInput = ""
	} else {
		shell.rawInput = cmmnd[:len(cmmnd)-1]

	}
	shell.command = cmd.Parse(shell.rawInput)
	shell.history = append(shell.history, cmmnd)

	return shell.command, nil
}

func (shell *Shell) eval() int8 {
	var cmnd cmd.Command = shell.Command()
	fnc, exist := blt.Builtins[cmnd.Base()]

	if exist {
		err := fnc(shell, cmnd)
		if err != nil {
			fmt.Print(err)
		}
	} else {
		if err := shell.executeExternalCommand(cmnd); err != nil {
			return shell.errorHandler(COMMAND_NOT_FOUND)
		}
	}
	return 0
}

func (shell *Shell) executeExternalCommand(cmnd cmd.Command) error {
	cmdPath := filepath.Join("/bin/", cmnd.Base())

	if _, err := os.Stat(cmdPath); os.IsNotExist(err) {
		return fmt.Errorf("command not found")
	}

	command := exec.Command(cmdPath, cmnd.Args()...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	if err := command.Run(); err != nil {
		return err
	}

	return nil
}

func (shell *Shell) errorHandler(err int8) int8 {
	var raw string = shell.command.Raw()
	if err == 1 {
		fmt.Fprintf(shell.config.Stdout(), "%s: %s: command not found", shell.name, raw[:len(raw)-1])
		return COMMAND_NOT_FOUND
	} else {
		fmt.Fprintf(shell.config.Stdout(), "%s: %s: command not found", shell.name, raw[:len(raw)-1])
	}
	return -1
}

func (shell *Shell) repl() {
	// var code int8
	for {
		shell.promptBuilder()
		shell.readInput()
		shell.eval()
	}
}

func Run() {
	var shell *Shell = New()
	shell.repl()
}
