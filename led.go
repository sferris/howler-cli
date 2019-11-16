package main

import (
  "fmt"
  "log"
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

/*
  if controller == nil {
    log.Printf("Opening howler device: %d\n", device)

    controller, err = howler.OpenDevice(device)
    if err != nil {
      log.Fatal(err.Error())
    }
  }
*/

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
  log.Printf("Setting %s LED color: %s\n", led.Name, led.Color);

  var ok bool
  var name  howler.Leds
  var rgb   color.RGBStruct

  name, ok = howler.Led(led.Name)
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

  err := controller.SetLEDRGB(name, rgb.Red, rgb.Green, rgb.Blue)
  if err != nil {
    return err
  }

  //result.Dump()

  return nil
}

func (led LEDStruct) setLEDDefault() error {
  log.Printf("Setting %s LED color: %s\n", led.Name, led.Color);

  var ok bool
  var name  howler.Leds
  var rgb   color.RGBStruct

  name, ok = howler.Led(led.Name)
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

  //log.Printf("%s\n", result.Dump())

  return nil
}
