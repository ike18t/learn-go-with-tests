package main

import "fmt"

func Hello(name string, language string) string {
  if name == "" {
    name = "World"
  }
  if language == "Spanish" {
    return "Hola, " + name
  }
  return "Hello, " + name
}

func main() {
  fmt.Println(Hello("world", "english"))
}
