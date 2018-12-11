package goinput

import (
	"strconv"
)

//Intin Take integer input with message mess
func Intin(mess string) (int, error) {
	text := Strin(mess)
	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}
	return num, nil
}
