package goinput

import (
	"strconv"
)

//Floatin take float64 input with message mess
func Floatin(mess string) (float64, error) {
	text := Strin(mess)
	num, ok := strconv.ParseFloat(text, 64)
	if ok != nil {
		return 0, ok
	}
	return num, nil
}
