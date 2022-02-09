package greetings

import (
  "fmt"
  "errors"
  "math/rand"
  "time"
)

func Hello(name string) (string, error) {
  if name == "" {
    return "", errors.New("ERR_NO_NAME")
  }
  message := fmt.Sprintf(randomFormat(), name)
  return message, nil
}

func randomFormat() string {
  formats := []string {
    "Hi, %v. Welcome!",
    "Nice to meet you, %v!",
    "Finally, we meet, %v!",
  }
  return formats[rand.Intn(len(formats))]
}

func init() {
  rand.Seed(time.Now().UnixNano())
}

