package cli

import (
	"errors"
	"reflect"
)

type String struct {
	Name     string
	Usage    string
	instance string
}

func (flagString *String) GetName() string {
	return flagString.Name
}

func (flagString *String) GetUsage() string {
	return flagString.Usage
}

func (flagString *String) Kind() reflect.Kind {
	return reflect.String
}

func (flagString *String) val() interface{} {
	return flagString.instance
}

func (flagString *String) set(args *ArgsArray) error {
	for i := 0; i < len(*args); i++ {
		if (*args)[i] == flagString.Name {
			if i+1 > len(*args) {
				return errors.New(flagString.Name + " needs argument")
			}
			flagString.instance = (*args)[i+1]
			args.remove(i, i+1)
			return nil
		}
	}
	return nil
}
