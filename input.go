package main

import (
  "fmt"
)

type InputStruct struct {
  Button    string `yaml:"button"`
  Mode      string `yaml:"mode"`
  Modifier  string `yaml:"modifier"`
  Value     string `yaml:"value"`
}

func (input *InputStruct) Process() error {
  fmt.Printf("Button: %s, Mode: %s, Modifier: %s, Value: %s\n\n",
    input.Button, input.Mode, input.Modifier, input.Value)
  return nil
}
