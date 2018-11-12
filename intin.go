package goinput

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Intin Take integer input with message mess
func Intin(mess string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(mess)
	text, _ := reader.ReadString('\n')
	num, ok := strconv.ParseInt(text[:len(text)-1], 10, 64)
	if ok != nil {
		//fmt.Println("That's not a number")
		return 0, ok
	}
	return int(num), nil
}
