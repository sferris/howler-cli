package main

import (
  "fmt"

  //howler "github.com/sferris/howler-controller"
)

type InputStruct struct {
  Command   string `yaml:"command"`

  Name      string `yaml:"input"`
  Type      string `yaml:"type"`
  Value     string `yaml:"value"`
  Modifier  string `yaml:"modifier"`
}

func (input *InputStruct) String() string {
  return fmt.Sprintf(
    "Input: %s, Type: %s, Modifier: %s, Value: %s",
      input.Name,
      input.Type,
      input.Modifier,
      input.Value)
}

func (input *InputStruct) Process() error {
  var err error

  switch input.Command {
    case "joystick-analog": fallthrough
    case "set-joystick-analog":
      err = input.SetJoystickAnalog()
    case "joystick-button": fallthrough
    case "set-joystick-button":
      err = input.SetJoystickButton()
    case "joystick-digital": fallthrough
    case "set-joystick-digital":
      err = input.SetJoystickDigital()
    case "keyboard-button": fallthrough
    case "set-keyboard-button":
      err = input.SetKeyboardButton()
    case "mouse-axis": fallthrough
    case "set-mouse-axis":
      err = input.SetMouseAxis()
    case "mouse-button": fallthrough
    case "set-mouse-button":
      err = input.SetMouseButton()
    default:
      return fmt.Errorf("Invalid command for processing: %s\n", input.Command)
  }

  return err
}

