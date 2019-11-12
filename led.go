package main

import (
  "fmt"

  "github.com/sferris/howler-controller/color"
)

type LedStruct struct {
  Button string `yaml:"button"`
  Scope  string `yaml:"scope"`
  Color  string `yaml:"color"`
}

func (led LedStruct) Process() error {
  if led.Scope == "" {
    led.Scope = "current"
  }

  if led.Button == "" {
    return fmt.Errorf("Mandtory --button option is missing")
  }

  rgb, ok := color.Lookup(led.Color); 
  if led.Color == "" {
    return fmt.Errorf("Mandtory color option is missing")
  } else if !ok {
    return fmt.Errorf("Invalid color given: %s", led.Color)
  }

  fmt.Printf("Button: %s, Color: %s, Scope: %s\n\n", led.Button, rgb.String(), led.Scope)
  return nil
}
