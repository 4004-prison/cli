package cli

import (
	"reflect"
)

type Bool struct {
	Name     string
	Usage    string
	instance bool
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
	return flagBool.instance
}

func (flagBool *Bool) set(args *ArgsArray) error {
	for i := 0; i < len(*args); i++ {
		if (*args)[i] == flagBool.Name {
			flagBool.instance = true
		}
	}
	return nil
}
