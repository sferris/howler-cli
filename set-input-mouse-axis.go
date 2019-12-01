package main

import (
  //"fmt"

  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)

/*
  Name: "input",
  --
  Name: "axis",
*/

func setMouseAxis(c *cli.Context) error {
  input := InputStruct{
    Name:      c.String(""),
    Type:      c.String(""),
    Modifier:  c.String(""),
    Value:     c.String(""),
  }

  return input.Process()
}

// SetMouseAxis(control ControlInput, function ControlFunction)
func (input *InputStruct) SetMouseAxis() error {
  var err error

  control, err := howler.StringToControl(input.Name)
  if err != nil {
    return err
  }

  function, err := howler.StringToControlFunction(input.Type)
  if err != nil {
    return err
  }

  if controller == nil { 
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return err
    }
  }   
    
  _, err = controller.SetMouseAxis(control, function)
  return err
}
