package main

import (
  "fmt"
  "strings"

  "github.com/sferris/howler-controller"
  "github.com/sferris/howler-controller/color"
)

type LEDStruct struct {
  Name   string `yaml:"led"`
  Scope  string `yaml:"scope"`
  Color  string `yaml:"color"`
}

func (led LEDStruct) Process() error {
  var err error

  if led.Scope == "" {
    led.Scope = "current"
  }

  if controller == nil {
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return err
    }
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
  var ok bool
  var name  howler.LedInputs
  var rgb   color.RGBStruct

  name = howler.ToLed(led.Name)
  if name == -1 {
    return fmt.Errorf(
      "Invalid LED Button reference: '%s': ",
      led.Name,
    )
  }
  rgb, ok = color.Lookup(led.Color); 
  if !ok {
    return fmt.Errorf("Invalid color value: %s", led.Color)
  }

  err := controller.SetLEDRGB(name, rgb.Red, rgb.Green, rgb.Blue)
  if err != nil {
    return err
  }

  //result.Dump()

  return nil
}

func (led LEDStruct) setLEDDefault() error {
  var ok bool
  var rgb   color.RGBStruct

  name := howler.ToLed(led.Name)
  if !ok {
    return fmt.Errorf(
      "Invalid LED Button reference: '%s': ",
      led.Name,
    )
  }
  rgb, ok = color.Lookup(led.Color); 
  if !ok {
    return fmt.Errorf("Invalid color value: %s", led.Color)
  }

  _, err := controller.SetDefaultLEDRGB(name, rgb.Red, rgb.Green, rgb.Blue)
  if err != nil {
    return err
  }

  return nil
}

func getLEDSettings() error {
  for led := howler.LedMin; led < howler.LedMax; led++ {
    var err error

    if controller == nil {
      controller, err = howler.OpenDevice(device)
      if err != nil {
        return err
      }
    }

    c, _ := controller.GetLEDColor(howler.LedInputs(led))
    fmt.Printf("Led %-15s %s\n", 
      fmt.Sprintf("%s:", howler.LedInputs(led)), c.ToIntString())
  }

  return nil
}
