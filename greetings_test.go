package greetings

import (
  "testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
  name := "Gladys"
  expected := "Hello, Gladys!"
  message, _ := Hello(name)

  if message != expected {
    t.Errorf("Hello(%v) != %v", name, expected)
  }
}


func TestHelloError(t *testing.T) {
  _, err := Hello("")

  if err == nil {
    t.Errorf("Hello doesn't raise an error with empty name")
  }
}
