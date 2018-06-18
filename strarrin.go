package goinput

import (
  "fmt"
  "bufio"
  "os"
)

func StrArrin(mess, sep string) ([]string){
  reader := bufio.NewReader(os.Stdin)
  fmt.Print(mess)
  text, _ := reader.ReadString('\n')
  var s []string
  c := 0
  for i := 0; i < len(text); i++{
    if string(text[i]) == sep{
      s = append(s, text[c:i])
      c = i+1
    }
  }
  s = append(s, text[c:len(text)-1])
  return s
}
