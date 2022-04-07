package greetings

import (
  "fmt"
  "errors"
)

func Hello(name string) (string, error) {
  if name == "" {
    return "", errors.New("Empty name")
  }
  message := fmt.Sprintf("Hello, %v!", name)
  return message, nil
}
