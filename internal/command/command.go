package command

type Command struct {
	base string
	// pragma []string
	args []string
}

func (c *Command) Base() string {
	return c.base
}

func (c *Command) Args() []string {
	return c.args
}
