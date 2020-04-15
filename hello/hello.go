package main

import "fmt"

var helloDict = map[string]string{"": "Hello", "French": "Bonjour", "Spanish": "Hola"}

func Hello(name string, language string) string {

  if name == "" {
    name = "World"
  }

  return helloDict[language] + ", " + name
}

func main() {
  fmt.Println(Hello("world", "english"))
}
