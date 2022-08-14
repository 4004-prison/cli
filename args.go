package cli

type ArgsArray []string

func (a *ArgsArray) First() string {
	return a.Get(0)
}

func (a *ArgsArray) Tail() []string {
	if len(*a) > 1 {
		ret := make([]string, len(*a)-1)
		copy(ret, (*a)[1:])
		return ret
	}
	return []string{}
}

func (a *ArgsArray) Get(index int) string {
	if len(*a) > index {
		return (*a)[index]
	}

	return ""
}

func (a *ArgsArray) remove(index ...int) {
	for _, i := range index {
		switch {
		case i == len(*a)-1:
			*a = (*a)[:i]
		case i >= 0 && i < len(*a)-1:
			*a = append((*a)[:i], (*a)[i+1:]...)
		}
	}
}
