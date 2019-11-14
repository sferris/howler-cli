package main

import (
  "fmt"
  "strings"

  "github.com/sferris/howler-controller"
  "github.com/sferris/howler-controller/color"
)

type LEDStruct struct {
  Button string `yaml:"button"`
  Scope  string `yaml:"scope"`
  Color  string `yaml:"color"`
}

func (led LEDStruct) Process() error {
  var err error

  if led.Scope == "" {
    led.Scope = "current"
  }

  switch strings.ToLower(led.Scope) {
    case "current":
      err = led.setLEDCurrent();

    case "default":
      err = led.setLEDDefault();

    default:
      return fmt.Errorf("Invalid LED Scope: %s\n", led.Scope)
  }

  return err
}

func (led LEDStruct) setLEDCurrent() error {
  fmt.Printf("Setting %s LED color: %s\n", led.Button, led.Color);

  var ok bool
  var button  howler.Leds
  var rgb     color.RGBStruct

  button, ok = howler.Button(led.Button)
  if !ok {
    return fmt.Errorf(
      "Invalid LED Button reference: '%s': ",
      led.Button,
    )
  }
  rgb, ok = color.Lookup(led.Color); 
  if !ok {
    return fmt.Errorf("Invalid color value: %s", led.Color)
  }

  err := controller.SetLEDRGB(button, rgb.Red, rgb.Green, rgb.Blue)
  if err != nil {
    return err
  }

  //result.Dump()

  return nil
}

func (led LEDStruct) setLEDDefault() error {
  fmt.Printf("Setting %s LED color: %s\n", led.Button, led.Color);

  var ok bool
  var button  howler.Leds
  var rgb     color.RGBStruct

  button, ok = howler.Button(led.Button)
  if !ok {
    return fmt.Errorf(
      "Invalid LED Button reference: '%s': ",
      led.Button,
    )
  }
  rgb, ok = color.Lookup(led.Color); 
  if !ok {
    return fmt.Errorf("Invalid color value: %s", led.Color)
  }

  result, err := controller.SetDefaultLEDRGB(button, rgb.Red, rgb.Green, rgb.Blue)
  if err != nil {
    return err
  }

  result.Dump()

  return nil
}
