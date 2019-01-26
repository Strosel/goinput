package goinput

import (
	"bufio"
	"fmt"
	"os"
)

//Strin take string input with message mess
func Strin(mess string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(mess)
	text, _ := reader.ReadString('\n')
	if len(text) < 1 {
		return ""
	} else if text[len(text)-2] == '\r' {
		return text[:len(text)-2]
	}
	return text[:len(text)-1]
}
