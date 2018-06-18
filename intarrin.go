package goinput

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func IntArrin(mess, sep string) ([]int){
  reader := bufio.NewReader(os.Stdin)
  fmt.Print(mess)
  text, _ := reader.ReadString('\n')
  var s []int
  c := 0
  for i := 0; i < len(text); i++{
    if string(text[i]) == sep{
      q, ok := strconv.ParseInt(text[c:i], 10, 64)
      if ok != nil {
        //ERROR noninteger
        c = i+1
      }else{
        s = append(s, int(q))
        c = i+1
      }
    }
  }
  q, ok := strconv.ParseInt(text[c:len(text)-1], 10, 64)
  if ok == nil{
    s = append(s, int(q))
  }
  return s
}
