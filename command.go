package cli

type ActionFunc func(*Context)

type Command struct {
	Name        string
	Usage       string
	Description string
	Action      ActionFunc
	Args        []Args
}

func (c *Command) Run(ctx *Context) {

}
