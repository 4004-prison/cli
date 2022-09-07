package cli

import (
	"errors"
	"reflect"
	"strconv"
)

type Int struct {
	Name  string
	Usage string
	Value int
}

func (flagInt *Int) GetName() string {
	return flagInt.Name
}

func (flagInt *Int) GetUsage() string {
	return flagInt.Usage
}

func (flagInt *Int) Kind() reflect.Kind {
	return reflect.String
}

func (flagInt *Int) val() interface{} {
	return flagInt.Value
}

func (flagInt *Int) set(args *ArgsArray) error {
	for i := 0; i < len(*args); i++ {
		if (*args)[i] == flagInt.Name {
			if i+1 > len(*args) {
				return errors.New(flagInt.Name + " needs argument")
			}
			instance, err := strconv.Atoi((*args)[i+1])
			if err != nil {
				return errors.New(flagInt.Name + " is not integer")
			}
			flagInt.Value = instance
			args.remove(i, i+1)
			return nil
		}
	}
	return nil
}

func (ctx *Context) Int(flag string) int {
	if any := ctx.Value(flag); any != nil {
		if v, ok := any.(int); ok {
			return v
		}
	}
	return 0
}
