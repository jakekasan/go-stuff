package main

import "fmt"

func main(){
  messages := make(chan string,2)

  messages<-"channel"
  messages<-"buffering"

  fmt.Println(<-messages)
  fmt.Println(<-messages)
}
