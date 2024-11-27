package main

import "fmt"


func sayHello(name string) string {
  return "hello world " + name
}


func main() {
  fmt.Println(sayHello("adem"))
}
