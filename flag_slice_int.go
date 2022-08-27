package cli

import (
	"errors"
	"reflect"
	"strconv"
)

type SliceInt struct {
	Name     string
	Usage    string
	instance []int
}

func (flagSliceInt *SliceInt) GetName() string {
	return flagSliceInt.Name
}

func (flagSliceInt *SliceInt) GetUsage() string {
	return flagSliceInt.Usage
}

func (flagSliceInt *SliceInt) Kind() reflect.Kind {
	return reflect.Slice
}

func (flagSliceInt *SliceInt) val() interface{} {
	return flagSliceInt.instance
}

func (flagSliceInt *SliceInt) set(args *ArgsArray) error {
	for i := 0; i < len(*args); i++ {
		if (*args)[i] == flagSliceInt.Name {
			if i+1 > len(*args) {
				return errors.New(flagSliceInt.Name + " needs argument")
			}
			instance, err := strconv.Atoi((*args)[i+1])
			if err != nil {
				return errors.New(flagSliceInt.Name + " is not integer")
			}
			flagSliceInt.instance = append(flagSliceInt.instance, instance)
			args.remove(i, i+1)
			i--
		}
	}
	return nil
}
