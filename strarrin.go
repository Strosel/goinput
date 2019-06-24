package goinput

import "strings"

//StrArrin Take string aray input with message mess, separated by sep
func StrArrin(mess, sep string) []string {
	text := Strin(mess)
	s := strings.Split(text, sep)
	return s
}

func formatter(sep, ex rune) func(rune) bool {
	exes := 0
	return func(r rune) bool {
		if r == ex {
			exes++
		} else if r == sep {
			return exes%2 == 0
		}
		return false
	}
}

//StrArrSpecin Take string aray input with message mess, separated by sep, exept if surrounded by ex
func StrArrSpecin(mess string, sep, ex rune) []string {
	if sep == ex {
		return StrArrin(string(mess), string(sep))
	}
	text := Strin(mess)
	// s := strings.Split(text, sep)
	s := strings.FieldsFunc(text, formatter(sep, ex))
	return s
}
