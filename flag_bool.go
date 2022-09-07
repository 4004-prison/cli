package cli

import (
	"reflect"
)

type Bool struct {
	Name  string
	Usage string
	Value bool
}

func (flagBool *Bool) GetName() string {
	return flagBool.Name
}

func (flagBool *Bool) GetUsage() string {
	return flagBool.Usage
}

func (flagBool *Bool) Kind() reflect.Kind {
	return reflect.String
}

func (flagBool *Bool) val() interface{} {
	return flagBool.Value
}

func (flagBool *Bool) set(args *ArgsArray) error {
	for i := 0; i < len(*args); i++ {
		if (*args)[i] == flagBool.Name {
			flagBool.Value = true
		}
	}
	return nil
}

func (ctx *Context) Bool(flag string) bool {
	if any := ctx.Value(flag); any != nil {
		if v, ok := any.(bool); ok {
			return v
		}
	}
	return false
}
