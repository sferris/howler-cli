package main

import (
  "fmt"
  "log"
  _"strings"

  howler "github.com/sferris/howler-controller"
)

type InputStruct struct {
  Name      string `yaml:"input"`
  Type      string `yaml:"type"`
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

  var ok bool
  var err error
  var inputType howler.InputTypes

  if inputType, ok = howler.ToInputType(input.Type); !ok {
    return fmt.Errorf("Invalid input type: %s\n", input.Type)
  }

  if controller == nil {
    log.Printf("Opening howler device: %d\n", device)

    controller, err = howler.OpenDevice(device)
    if err != nil {
      log.Fatal(err.Error())
    }
  }
 
  switch inputType {
    case howler.TypeJoystick1: fallthrough
    case howler.TypeJoystick2:
      err = input.setJoystickInput();

    case howler.TypeKeyboard:
      err = input.setKeyboardInput();

    case howler.TypeMouse:
      err = input.setMouseInput();
  }

  return err
}

func (input *InputStruct) setKeyboardInput() error {
  log.Printf("Setting keyboard: %+v\n\n", input);

  var ok bool
  var name  howler.Inputs
  var key   howler.Keys
  var mod   howler.Modifiers

  name, ok = howler.Input(input.Name)
  if !ok {
    return fmt.Errorf(
      "Invalid input name: '%s': ",
      input.Name,
    )
  }
  key, ok = howler.ToKey(input.Value)
  if !ok {
    return fmt.Errorf(
      "Invalid keyboard code for input: '%s': ",
      input.Value,
    )
  }
  mod, ok = howler.ToModifier(input.Modifier)
  if !ok {
    return fmt.Errorf(
      "Invalid keyboard modifier for input: '%s': ",
      input.Modifier,
    )
  }

  result, err := controller.SetInputKeyboard(name, key, mod)
  if err != nil {
    return err
  }

  result.Dump()

  return nil
}

func (input *InputStruct) setJoystickInput() error {
  log.Println("Setting joystick");

  var ok bool
  var name      howler.Inputs
  var joystick  howler.InputTypes
  var button    howler.JoystickButtons

  name, ok = howler.Input(input.Name)
  if !ok {
    return fmt.Errorf(
      "Invalid input name: '%s': ",
      input.Name,
    )
  }
  joystick, ok = howler.ToInputType(input.Type)
  if !ok {
    return fmt.Errorf(
      "Invalid input joystick: '%s': ",
      input.Type,
    )
  }
  button, ok = howler.ToJoystickButton(input.Value)
  if !ok {
    return fmt.Errorf(
      "Invalid joystick button number for input: '%s': (Must be 1-32)", 
      input.Value,
    )
  }

  result, err := controller.SetInputJoystick(name, joystick, button)
  if err != nil {
    return err
  }

  result.Dump()

  return nil
}

func (input *InputStruct) setMouseInput() error {
  log.Println("Setting mouse");

  var ok bool
  var name   howler.Inputs
  var button howler.MouseButtons

  name, ok = howler.Input(input.Name)
  if !ok {
    return fmt.Errorf(
      "Invalid input name: '%s': ",
      input.Name,
    )
  }
  button, ok = howler.ToMouseButton(input.Value)
  if !ok {
    return fmt.Errorf(
      "Invalid mouse button number for input: '%s': ",
      input.Value,
    )
  }

  result, err := controller.SetInputMouse(name, button)
  if err != nil {
    return err
  }

  result.Dump()

  return nil
}
