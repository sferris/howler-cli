package main

import (
  "fmt"
  "io/ioutil"

  "gopkg.in/yaml.v2"

  //"github.com/sferris/howler-controller"
)

type FileStruct struct {
  Path    string

  Game    string                 `yaml: "game"`
  Colors  *map[string]RGBStruct  `yaml:"colors,omitempty"`
  Leds    []LedStruct            `yaml: "leds"`
  Inputs  []InputStruct          `yaml: "inputs"`

}

func (file FileStruct) Process() error {
  file.Colors = &colors;

  fmt.Printf("Filename: %s\n\n", file.Path)

  source, err := ioutil.ReadFile(file.Path)
  if err != nil {
    return err
  }

  err = yaml.Unmarshal(source, &file)
  if err != nil {
      return err
  }

  for _, input := range file.Inputs {
    input.Process();
  }

  for _, led := range file.Leds {
    led.Process();
  }

  return nil
}
