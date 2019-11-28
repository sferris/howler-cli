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

  control := howler.ToControl(input.Name)
  if len(control.Type()) <= 0 {
    return fmt.Errorf(
      "Invalid input name: '%s': ",
      input.Name,
    )
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

  _, err := controller.SetInputKeyboard(control.Input(), key, mod)
  if err != nil {
    return err
  }

  return nil
}

func (input *InputStruct) setButtonJoystick() error {

fmt.Println(input)

  // default joystick
  if input.Type == "" {
    input.Type = "joystick1"
  }

  control := howler.ToControl(input.Name)
  if len(control.Type()) <= 0 {
    return fmt.Errorf(
      "Invalid input name: '%s': ",
      input.Name,
    )
  }
  joystick := howler.ToInputType(input.Type)
  if joystick == -1 {
    return fmt.Errorf(
      "Invalid input joystick: '%s': ",
      input.Type,
    )
  }
  button := howler.ToJoystickButton(input.Value)
  if button == -1 {
    return fmt.Errorf(
      "Invalid joystick button number for input: '%s'",
      input.Value,
    )
  }

  _, err := controller.SetInputJoystick(control.Input(), joystick, button)
  if err != nil {
    return err
  }

  return nil
}

func (input *InputStruct) setButtonMouse() error {
  control := howler.ToControl(input.Name)
  if len(control.Type()) <= 0 {
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

  _, err := controller.SetInputMouse(control.Input(), button)
  if err != nil {
    return err
  }

  return nil
}
