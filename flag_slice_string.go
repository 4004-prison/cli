package cli

import (
	"errors"
	"reflect"
)

type SliceString struct {
	Name     string
	Usage    string
	instance []string
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
	return flagSliceString.instance
}

func (flagSliceString *SliceString) set(args ArgsArray) error {
	// for i, v := range args {
	// if v == flagString.Name {
	// 	if i+1 > len(args) {
	// 		return errors.New(flagString.Name + " needs argument")
	// 	}
	// 	flagString.instance = args[i+1]
	// 	args.remove(i, i+1)
	// }
	// }

	for i := 0; i < len(args); i++ {
		if args[i] == flagSliceString.Name {
			if i+1 > len(args) {
				return errors.New(flagSliceString.Name + " needs argument")
			}
			flagSliceString.instance = append(flagSliceString.instance, args[i+1])
			args.remove(i, i+1)
			i--
			// TODO:
		}
	}
	return nil
}
