package main

import (
  "log"
  "fmt"
  "strings"
  "strconv"

  howler "github.com/sferris/howler-controller"
)

type rgbFlags struct {
  Red   int
  Blue  int
  Green int
}

func (rgb *rgbFlags) String() string {
  return "foo"
}

func (rgb *rgbFlags) Set(value string) error {
  parsed := strings.Split(value, ",")

  var red, blue, green int
  var err error

  if red, err = strconv.Atoi(parsed[0]); err != nil {
    log.Fatalf("Invalid red value: %s\n", parsed[0])
  }
  if green, err = strconv.Atoi(parsed[1]); err != nil {
    log.Fatalf("Invalid green value: %s\n", parsed[1])
  }
  if blue, err = strconv.Atoi(parsed[2]); err != nil {
    log.Fatalf("Invalid blue value: %s\n", parsed[2])
  }
  *rgb = rgbFlags{red, green, blue}

  return nil
}

func (led *Led) String() string {
  return fmt.Sprintf(
    "Button: %d, Scope: %s, Red: %d, Green: %d, Blue: %d",
      howler.Button(led.Button),
      led.Scope,
      led.RGB.Red,
      led.RGB.Green,
      led.RGB.Blue,
  )
}

