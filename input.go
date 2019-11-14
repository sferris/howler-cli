package main

import (
  "fmt"
  "strings"
  _"strconv"

  howler "github.com/sferris/howler-controller"
)

type InputStruct struct {
  Name      string `yaml:"input"`
  Mode      string `yaml:"mode"`
  Modifier  string `yaml:"modifier"`
  Value     string `yaml:"value"`
}

func (input *InputStruct) Process() error {
  if input.Modifier == "" {
    input.Modifier = "none"
  }
  if input.Name == "" {
    return fmt.Errorf("Mandatory input option is missing")
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

  var ok bool
  var name  howler.Inputs
  var mode  howler.Modes
  var value howler.InputValues
  mod := howler.ModifierNone

  name, ok = howler.Input(input.Name)
  if !ok {
    return fmt.Errorf(
      "Invalid input name: '%s': ",
      input.Name,
    )
  }
  mode, ok = howler.Mode(input.Mode)
  if !ok {
    return fmt.Errorf(
      "Invalid input mode: '%s': ",
      input.Mode,
    )
  }
  value, ok = howler.JoystickButton(input.Value)
  if !ok {
    return fmt.Errorf(
      "Invalid joystick button number for input: '%s': (Must be 1-32)", 
      input.Value,
    )
  }

  result, err := controller.SetInput(name, mode, value, mod)
  if err != nil {
    return err
  }

  result.Dump()

  return nil
}

func (input *InputStruct) setKeyboardInput() error {
  fmt.Printf("Setting keyboard: %+v\n\n", input);

  var ok bool
  var name  howler.Inputs
  var mode  howler.Modes
  var value howler.InputValues
  var mod   howler.Modifiers

  name, ok = howler.Input(input.Name)
  if !ok {
    return fmt.Errorf(
      "Invalid input name: '%s': ",
      input.Name,
    )
  }
  mode, ok = howler.Mode(input.Mode)
  if !ok {
    return fmt.Errorf(
      "Invalid input mode: '%s': ",
      input.Mode,
    )
  }
  value, ok = howler.Key(input.Value)
  if !ok {
    return fmt.Errorf(
      "Invalid keyboard code for input: '%s': ",
      input.Value,
    )
  }
  mod, ok = howler.Modifier(input.Modifier)
  if !ok {
    return fmt.Errorf(
      "Invalid keyboard modifier for input: '%s': ",
      input.Modifier,
    )
  }

  result, err := controller.SetInput(name, mode, value, mod)
  if err != nil {
    return err
  }

  result.Dump()

  return nil
}
func (input *InputStruct) setMouseInput() error {
  fmt.Println("Setting mouse");

  var ok bool
  var name  howler.Inputs
  var mode  howler.Modes
  var value howler.InputValues
  mod := howler.ModifierNone

  name, ok = howler.Input(input.Name)
  if !ok {
    return fmt.Errorf(
      "Invalid input name: '%s': ",
      input.Name,
    )
  }
  mode, ok = howler.Mode(input.Mode)
  if !ok {
    return fmt.Errorf(
      "Invalid input mode: '%s': ",
      input.Mode,
    )
  }
  value, ok = howler.MouseButton(input.Value)
  if !ok {
    return fmt.Errorf(
      "Invalid mouse button number for input: '%s': ",
      input.Value,
    )
  }

  result, err := controller.SetInput(name, mode, value, mod)
  if err != nil {
    return err
  }

  result.Dump()

  return nil
}
