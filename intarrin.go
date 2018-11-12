package goinput

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//IntArrin Take integer array input with message mess, separated by sep
func IntArrin(mess, sep string) []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(mess)
	text, _ := reader.ReadString('\n')
	var s []int
	c := 0
	for i := 0; i < len(text); i++ {
		if string(text[i]) == sep {
			q, ok := strconv.ParseInt(text[c:i], 10, 64)
			if ok != nil {
				c = i + 1
			} else {
				s = append(s, int(q))
				c = i + 1
			}
		}
	}
	q, ok := strconv.ParseInt(text[c:len(text)-1], 10, 64)
	if ok == nil {
		s = append(s, int(q))
	}
	return s
}
