package main

import (
  "fmt"

  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)

/*
  Name: "input",
  --
  Name: "function",
*/

var txtJoystickAnalog = `
This command will set a controller input to emit analog joystick movements.

Inputs: %s

Functions: %s
`

var usageJoystickAnalog = "Set input to emit analog joystick movements"
var descJoystickAnalog = fmt.Sprintf(
  txtJoystickAnalog,
  fmt.Sprintf("\n%s", ControlInputNameByCapability(howler.CapJoystickAnalog)),
  fmt.Sprintf("\n%s", ControlFunctionNameByCapability(howler.CapJoystickAnalog)),
)

func setJoystickAnalog(c *cli.Context) error {
  input := InputStruct{
    Command:   "set-joystick-analog",

    Name:      c.String("input"),
    Type:      c.String("function"),
    //Value:     c.String(""),
    //Modifier:  c.String(""),
  }

  return input.Process()
}

// SetJoystickAnalog(control ControlInput, function ControlFunction)
func (input *InputStruct) SetJoystickAnalog() error {
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
    
  _, err = controller.SetJoystickAnalog(control, function)
  return err
}
