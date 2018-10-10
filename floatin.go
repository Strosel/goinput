package goinput

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func Floatin(mess string) (float64, error){
  reader := bufio.NewReader(os.Stdin)
  fmt.Print(mess)
  text, _ := reader.ReadString('\n')
  num, ok := strconv.ParseFloat(text[:len(text)-1], 10, 64)
  if ok != nil {
    //fmt.Println("That's not a number")
    return 0, ok
  }else{
    return num, nil
  }
}
