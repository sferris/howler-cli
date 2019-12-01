package main

import (
  //"fmt"

  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)


/*
  Name: "input",
  Name: "function",
  Name: "value",
*/

func setJoystickDigital(c *cli.Context) error {
  input := InputStruct{
    Name:      c.String(""),
    Type:      c.String(""),
    Modifier:  c.String(""),
    Value:     c.String(""),
  }

  return input.Process()
}

// SetJoystickDigital(control ControlInput, function ControlFunction, value int8)
func (input *InputStruct) SetJoystickDigital() error {
  var err error

  control, err := howler.StringToControl(input.Name)
  if err != nil {
    return err
  }

  function, err := howler.StringToControlFunction(input.Type)
  if err != nil {
    return err
  }

  value := int8(1)

  if controller == nil { 
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return err
    }
  }   
    
  _, err = controller.SetJoystickDigital(control, function, value)
  return err
}
