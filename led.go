package main

import (
  "fmt"
  //"github.com/sferris/howler-controller/color"
)

type LedStruct struct {
  Button string `yaml:"button"`
  Scope  string `yaml:"scope"`
  Color  string `yaml:"color"`
}

func (led LedStruct) Process() error {
  if rgb, ok := FetchRGB(led.Color); ok {
    fmt.Printf("Button: %s, Color: %s, Scope: %s\n\n", led.Button, rgb.String(), led.Scope)
  }
  return nil
}
