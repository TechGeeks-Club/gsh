package builtins

import (
	"fmt"
	"gsh/internal/command"
)

func Help(c command.Command) {
	fmt.Print("gsh is a lightweight, custom shell implemented from scratch in Go. It provides a simple REPL to execute commands, handle input/output, and spawn processes, aiming to give developers a clear understanding of shell internals while offering a functional CLI environment.")
}
