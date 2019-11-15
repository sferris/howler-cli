package main

import (
  "log"
  "io/ioutil"

  "gopkg.in/yaml.v2"

  "github.com/sferris/howler-controller/color"
)


type FileStruct struct {
  Path    string

  Game    string                       `yaml: "game"`
  Colors  *map[string]color.RGBStruct  `yaml: "colors,omitempty"`
  Leds    []LEDStruct                  `yaml: "leds"`
  Inputs  []InputStruct                `yaml: "inputs"`
}

func (file FileStruct) Process() error {
  file.Colors = &colors;

  source, err := ioutil.ReadFile(file.Path)
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
  }

  return nil
}
