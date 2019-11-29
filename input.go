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

  typ, err := howler.StringToControlFunction(input.Type)
  if err != nil {
    return err
  }
 
  if typ.Capability() & howler.CapJoystickButton != 0 {
    return input.setButtonJoystick();
  } else if typ.Capability() & howler.CapKeyboardButton != 0 {
    return input.setButtonKeyboard();
  } else if typ.Capability() & howler.CapMouseButton != 0 {
    return input.setButtonMouse();
  }

  return nil
}

func getControlSettings() error {
  var err error

  if controller == nil {
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return err
    }
  }

/*
  for k, v := range howler.ControlFunctionMap {
    fmt.Printf("%d: %+v\n", k, v)
  }
*/

  for input := range howler.ControlInputs() {
    //input := howler.ControlButton26.ID()
    i, _ := controller.GetInput(howler.ControlID(input))
    fmt.Printf("%d: %s\n", input, i.String())
  }

  return nil
}
