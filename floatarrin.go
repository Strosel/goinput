package goinput

import (
	"errors"
	"fmt"
	"strconv"
)

//FloatArrin Take float64 array input with message mess, separated by sep. Non floats are ignored
func FloatArrin(mess, sep string) []float64 {
	if sep == "." {
		panic(errors.New("Separator can't be \".\""))
	}
	floats := []float64{}
	strings := StrArrin(mess, sep)
	for _, s := range strings {
		f, err := strconv.ParseFloat(s, 64)
		if err == nil {
			floats = append(floats, f)
		}
	}
	return floats
}

//FloatArrinE Take float64 array input with message mess, separated by sep. Non floats trigger error return
func FloatArrinE(mess, sep string) ([]float64, error) {
	if sep == "." {
		panic(errors.New("Separator can't be \".\""))
	}
	floats := []float64{}
	strings := StrArrin(mess, sep)
	for x, s := range strings {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return []float64{}, fmt.Errorf("Non float at index %v", x)
		}
		floats = append(floats, f)
	}
	return floats, nil
}
