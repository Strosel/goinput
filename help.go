package goinput

import (
  "fmt"
)

func Help()  {
  fmt.Println("String input - Strin(mess string) (string)")
  fmt.Println("Integer input - Intin(mess string) (int, error)")
  fmt.Println("String Array input - StrArrin(mess, sep string) ([]string)")
  fmt.Println("Integer Array input - IntArrin(mess, sep string) ([]int)")
}
