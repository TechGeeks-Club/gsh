package command

type Command struct {
	raw  string
	base string
	// pragma []string
	args []string
}

func (c *Command) Base() string {
	return c.base
}

func (c *Command) Raw() string {
	return c.raw
}
func (c *Command) Args() []string {
	return c.args
}
