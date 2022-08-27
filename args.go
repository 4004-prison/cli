package cli

import "sort"

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
	index = sort.IntSlice(index)
	indexCursor := 0
	newSlice := (*a)[:0]
	for i := range *a {
		if indexCursor < len(index) && i == index[indexCursor] {
			indexCursor++
			continue
		}
		newSlice = append(newSlice, (*a)[i])
	}
	*a = newSlice
}
