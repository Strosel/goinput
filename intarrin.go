package goinput

import (
	"fmt"
	"strconv"
)

//IntArrin Take integer array input with message mess, separated by sep. Non integers are ignored
func IntArrin(mess, sep string) []int {
	ints := []int{}
	strings := StrArrin(mess, sep)
	for _, s := range strings {
		i, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

//IntArrinE Take integer array input with message mess, separated by sep. Non integers trigger error return
func IntArrinE(mess, sep string) ([]int, error) {
	ints := []int{}
	strings := StrArrin(mess, sep)
	for x, s := range strings {
		i, err := strconv.Atoi(s)
		if err != nil {
			return []int{}, fmt.Errorf("Non integer at index %v", x)
		}
		ints = append(ints, i)
	}
	return ints, nil
}
