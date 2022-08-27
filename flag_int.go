package cli

import (
	"errors"
	"reflect"
	"strconv"
)

type Int struct {
	Name     string
	Usage    string
	instance int
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
	return flagInt.instance
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
			flagInt.instance = instance
			args.remove(i, i+1)
			return nil
		}
	}
	return nil
}
