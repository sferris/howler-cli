package main

import (
  "log"
  "io/ioutil"

  "time"

  "gopkg.in/yaml.v2"
  "gopkg.in/urfave/cli.v2"

  "github.com/sferris/howler-controller/color"
)

var colors = color.ColorMap

type FileStruct struct {
  Path    string

  Game    string                       `yaml: "game"`
  Colors  *map[string]color.RGBStruct  `yaml: "colors,omitempty"`
  Leds    []LEDStruct                  `yaml: "leds"`
  Inputs  []InputStruct                `yaml: "inputs"`
}

func readFile(c *cli.Context) error {
  file := FileStruct {
    Colors:   &colors,
  }

  source, err := ioutil.ReadFile(c.String("path"))
  if err != nil {
    return err
  }

  err = yaml.Unmarshal(source, &file)
  if err != nil {
    return err
  }

  for _, input := range file.Inputs {
    err := input.Process();
    if err != nil {
      log.Println(err.Error())
    }
  }

  for _, led := range file.Leds {
    err := led.Process();
    if err != nil {
      log.Println(err.Error())
    }

    if sleep := c.Int("sleep"); sleep > 0 {
      time.Sleep(time.Millisecond*time.Duration(sleep))
    }
  }

  return nil
}
