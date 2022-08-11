package cli

type Context struct {
	args []string
}

func ctx() *Context {
	return &Context{}
}
