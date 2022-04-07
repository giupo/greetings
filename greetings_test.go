package greetings

import (
  "testing"
  "strings"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloContainsName(t *testing.T) {
  name := "Gladys"
  message, _ := Hello(name)

  if !strings.Contains(message, name) {
    t.Errorf("Hello(%v) doesn't contain", name)
  }
}


func TestHelloError(t *testing.T) {
  _, err := Hello("")

  if err == nil {
    t.Errorf("Hello doesn't raise an error with empty name")
  }
}
