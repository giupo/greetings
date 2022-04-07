package greetings

import (
  "fmt"
  "errors"
  "math/rand"
  "time"
  "log"
)

func Hello(name string) (string, error) {
  if name == "" {
    return "", errors.New("Empty name")
  }

  message := fmt.Sprintf(randomFormat(), name)
  return message, nil
}


func init() {
  log.SetPrefix("greetings: ")
  log.SetFlags(0)
  rand.Seed(time.Now().UnixNano())
  log.Println("init")
}


func randomFormat() string {
  formats := []string{
    "Hi, %v. Welcome!",
    "Great to see you, %v!",
    "Hail, %v! Well met!",
  }

  return formats[rand.Intn(len(formats))]
}
