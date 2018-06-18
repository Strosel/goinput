package goinput

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func Intin(mess string) (int, error){
  reader := bufio.NewReader(os.Stdin)
  fmt.Print(mess)
  text, _ := reader.ReadString('\n')
  num, ok := strconv.ParseInt(text[:len(text)-1], 10, 64)
  if ok != nil {
    //fmt.Println("That's not a number")
    return 0, ok
  }else{
    return int(num), nil
  }
}
