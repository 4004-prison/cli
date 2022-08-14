package cli

type ActionFunc func(*Context)

type Command struct {
	Name        string
	Usage       string
	Description string
	Action      ActionFunc
	Flags       FlagSet
}

func (c *Command) Run(ctx *Context) {
	c.Action(ctx)
}
