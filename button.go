package main

import (
  "fmt"
  _"strings"

  howler "github.com/sferris/howler-controller"
)

func (input *InputStruct) setButtonKeyboard() error {
  // defaults
  if input.Modifier == "" {
    input.Modifier = "none"
  }

  control, err := howler.StringToControl(input.Name)
  if err != nil {
    return err
  }
  key := howler.ToKey(input.Value)
  if key == -1 {
    return fmt.Errorf(
      "Invalid keyboard code for input: '%s': ",
      input.Value,
    )
  }
  mod := howler.ToModifier(input.Modifier)
  if mod == -1 {
    return fmt.Errorf(
      "Invalid keyboard modifier for input: '%s': ",
      input.Modifier,
    )
  }

  {
    _, err := controller.SetInputKeyboard(control.ID(), key, mod)
    if err != nil {
      return err
    }
  }

  return nil
}

func (input *InputStruct) setButtonJoystick() error {

fmt.Println(input)

  // default joystick
  if input.Type == "" {
    input.Type = "joystick1"
  }

  control, err := howler.StringToControl(input.Name)
  if err != nil {
    return err
  }
  joystick, err := howler.StringToControlFunction(input.Type)
  if err != nil {
    return err
  }
  button := howler.ToJoystickButton(input.Value)
  if button == -1 {
    return fmt.Errorf(
      "Invalid joystick button number for input: '%s'",
      input.Value,
    )
  }

  {
    _, err := controller.SetInputJoystick(control.ID(), joystick.ID(), button)
    if err != nil {
      return err
    }
  }

  return nil
}

func (input *InputStruct) setButtonMouse() error {
  control, err := howler.StringToControl(input.Name)
  if err != nil {
    return fmt.Errorf(
      "Invalid input name: '%s': ",
      input.Name,
    )
  }
  button := howler.ToMouseButton(input.Value)
  if button == -1 {
    return fmt.Errorf(
      "Invalid mouse button number for input: '%s': ",
      input.Value,
    )
  }

  {
    _, err := controller.SetInputMouse(control.ID(), button)
    if err != nil {
      return err
    }
  }

  return nil
}
