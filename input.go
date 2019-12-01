package main

import (
  "fmt"

  //"strings"

  //howler "github.com/sferris/howler-controller"
)

type InputStruct struct {
  Name      string `yaml:"input"`
  Type      string `yaml:"type"`
  Modifier  string `yaml:"modifier"`
  Value     string `yaml:"value"`
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

  switch input.Type {
    case "joystick-analog":
      err = input.SetJoystickAnalog()
    case "joystick-button":
      err = input.SetJoystickButton()
    case "joystick-digital":
      err = input.SetJoystickDigital()
    case "keyboard-button":
      err = input.SetKeyboardButton()
    case "mouse-axis":
      err = input.SetMouseAxis()
    case "mouse-button":
      err = input.SetMouseButton()
    default:
      return fmt.Errorf("Invalid Input setting: %s\n", input.Type)
  }

  return err
}
