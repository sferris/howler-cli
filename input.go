package main

import (
  "fmt"
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
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return err
    }
  }

  inputType := howler.ToInputType(input.Type)
 
  switch inputType {
    case howler.TypeJoystick1: fallthrough
    case howler.TypeJoystick2:
      err = input.setButtonJoystick();

    case howler.TypeKeyboard:
      err = input.setButtonKeyboard();

    case howler.TypeMouse:
      err = input.setButtonMouse();
  }

  return err
}

func getControlSettings() error {
  for input := howler.ControlMin; input < howler.ControlMax; input++ {
    var err error

    if controller == nil {
      controller, err = howler.OpenDevice(device)
      if err != nil {
        return err
      }
    }

    i, _ := controller.GetInput(howler.ControlInput(input))

    fmt.Printf("%s\n", i.String())
/*
    fmt.Printf("Input %-15s %-25s %-3d %-3d\n", 
      fmt.Sprintf("%s:", howler.Inputs(input)),
      howler.InputTypes(i.InputType),
      i.InputValue1,
      i.InputValue2,
    )
*/
  }

  return nil
}
