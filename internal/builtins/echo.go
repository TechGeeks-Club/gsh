package builtins

import (
	"fmt"
	cmd "gsh/internal/command"
)

func Echo(c cmd.Command) {
	fmt.Println(c.Raw()[5:])
}
