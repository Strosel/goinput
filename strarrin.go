package goinput

import "strings"

//StrArrin Take string aray input with message mess, separated by sep
func StrArrin(mess, sep string) []string {
	text := Strin(mess)
	s := strings.Split(text, sep)
	return s
}
