package goinput

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Floatin take float64 input with message mess
func Floatin(mess string) (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(mess)
	text, _ := reader.ReadString('\n')
	num, ok := strconv.ParseFloat(text[:len(text)-1], 64)
	if ok != nil {
		return 0, ok
	}
	return num, nil
}
