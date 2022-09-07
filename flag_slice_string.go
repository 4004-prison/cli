package cli

import (
	"errors"
	"reflect"
)

type SliceString struct {
	Name  string
	Usage string
	Value []string
}

func (flagSliceString *SliceString) GetName() string {
	return flagSliceString.Name
}

func (flagSliceString *SliceString) GetUsage() string {
	return flagSliceString.Usage
}

func (flagSliceString *SliceString) Kind() reflect.Kind {
	return reflect.Slice
}

func (flagSliceString *SliceString) val() interface{} {
	return flagSliceString.Value
}

func (flagSliceString *SliceString) set(args *ArgsArray) error {
	for i := 0; i < len(*args); i++ {
		if (*args)[i] == flagSliceString.Name {
			if i+1 > len(*args) {
				return errors.New(flagSliceString.Name + " needs argument")
			}
			flagSliceString.Value = append(flagSliceString.Value, (*args)[i+1])
			args.remove(i, i+1)
			i--
		}
	}
	return nil
}

func (ctx *Context) SliceString(flag string) []string {
	if any := ctx.Value(flag); any != nil {
		if v, ok := any.([]string); ok {
			return v
		}
	}
	return nil
}
