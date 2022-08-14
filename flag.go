package cli

import "reflect"

type Flag interface {
	GetName() string
	GetUsage() string
	Kind() reflect.Kind
	val() interface{}
	set(ArgsArray) error
}

type FlagSet []Flag
