package main

import (
  "fmt"
  "strings"
)

type InputStruct struct {
  Input     string `yaml:"input"`
  Mode      string `yaml:"mode"`
  Modifier  string `yaml:"modifier"`
  Value     string `yaml:"value"`
}

func (input *InputStruct) Process() error {
  if input.Modifier == "" {
    input.Modifier = "none"
  }
  if input.Input == "" {
    return fmt.Errorf("Mandatory button option is missing")
  }

  var err error

  switch strings.ToLower(input.Mode) {
    case "joystick1": fallthrough
    case "joystick2":
      err = input.setJoystickInput();

    case "keyboard":
      err = input.setKeyboardInput();

    case "mouse":
      err = input.setMouseInput();

    default:
      return fmt.Errorf("Invalid input mode: %s\n", input.Mode)
  }

  return err
}

func (input *InputStruct) setJoystickInput() error {
  fmt.Println("Setting joystick");
  return nil
}
func (input *InputStruct) setKeyboardInput() error {
  fmt.Println("Setting keyboard");
  return nil
}
func (input *InputStruct) setMouseInput() error {
  fmt.Println("Setting mouse");
  return nil
}
