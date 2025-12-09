package builtins

import "gsh/internal/command"

var Builtins = map[string]func(c command.Command){
	"echo": Echo,
	"help": Help,
	"exit": Exit,
}
