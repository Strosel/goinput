package goinput

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

//FloatArrin Take float64 array input with message mess, separated by sep
func FloatArrin(mess, sep string) ([]float64, error) {
	if sep == "." {
		return []float64{}, errors.New("Separator can't be \".\"")
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(mess)
	text, _ := reader.ReadString('\n')
	var s []float64
	c := 0
	for i := 0; i < len(text); i++ {
		if string(text[i]) == sep {
			q, ok := strconv.ParseFloat(text[c:i], 64)
			if ok != nil {
				c = i + 1
			} else {
				s = append(s, q)
				c = i + 1
			}
		}
	}
	q, ok := strconv.ParseFloat(text[c:len(text)-1], 64)
	if ok == nil {
		s = append(s, q)
	}
	return s, nil
}
