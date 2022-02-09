package greetings

import (
  "fmt"
  "errors"
)

func Hello(name string) (string, error) {
  if name == "" {
    return "", errors.New("ERR_NO_NAME")
  }
  message := fmt.Sprintf("Hi, %v. Welcome!", name)
  return message, nil
}

