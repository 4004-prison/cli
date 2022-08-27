package cli

import (
	"log"
)

type Context struct {
	args    ArgsArray
	flagVal map[string]interface{}
}

func initCTX(args ArgsArray) *Context {
	if len(args) == 1 {
		log.Println("need a command")
		return nil
	}

	return &Context{
		args: args[1:],
	}
}

func (ctx *Context) getCMD() string {
	return ctx.args.First()
}

func (ctx *Context) Args() ArgsArray {
	return ctx.args
}

func (ctx *Context) parse(flagSet FlagSet) error {
	ctx.flagVal = make(map[string]interface{})
	iterator := ArgsArray(ctx.args.Tail())
	for _, flag := range flagSet {
		if err := flag.set(&iterator); err != nil {
			return err
		}
		ctx.flagVal[flag.GetName()] = flag.val()
	}
	return nil
}

func (ctx *Context) String(flag string) string {
	if any := ctx.Value(flag); any != nil {
		if v, ok := any.(string); ok {
			return v
		}
	}
	return ""
}

func (ctx *Context) Value(flag string) interface{} {
	if any, ok := ctx.flagVal[flag]; ok {
		return any
	}
	return nil
}
