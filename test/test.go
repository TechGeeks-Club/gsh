package main

import (
	"fmt"
	"gsh/internal/config"
)

func main() {
	var c config.Config
	c.NewConfig(nil, nil, nil, nil, nil, nil, nil)

	fmt.Print(c.Prompt)
}
