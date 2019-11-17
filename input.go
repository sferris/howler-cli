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

  if controller == nil {
    log.Printf("Opening howler device: %d\n", device)

    controller, err = howler.OpenDevice(device)
    if err != nil {
      log.Fatal(err.Error())
    }
  }

  inputType := howler.ToInputType(input.Type)
 
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
  log.Printf("Setting keyboard: %s\n", input);

  // defaults
  if input.Modifier == "" {
    input.Modifier = "none"
  }

  name := howler.ToInput(input.Name)
  if name == -1 {
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

  result, err := controller.SetInputKeyboard(name, key, mod)
  if err != nil {
    return err
  }

  result.Dump()

  return nil
}

func (input *InputStruct) setJoystickInput() error {
  log.Printf("Setting joystick: %s\n", input);

  // default joystick
  if input.Name == "" {
    input.Name = "joystick1"
  }

  name := howler.ToInput(input.Name)
  if name == -1 {
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

  result, err := controller.SetInputJoystick(name, joystick, button)
  if err != nil {
    return err
  }

  result.Dump()

  return nil
}

func (input *InputStruct) setMouseInput() error {
  log.Printf("Setting mouse: %s\n", input);

  name := howler.ToInput(input.Name)
  if name == -1 {
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

  result, err := controller.SetInputMouse(name, button)
  if err != nil {
    return err
  }

  result.Dump()

  return nil
}
