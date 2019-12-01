package main

import (
  "fmt"

  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)

/*
  Name: "input",
  --
  Name: "axis",
*/


var txtMouseAxis = `
This command will set a controller input to emit a mouse movements

Inputs: %s

Axis: %s
`

var usageMouseAxis = "Set input to emit mouse movements"
var descMouseAxis = fmt.Sprintf(
  txtMouseAxis,
  fmt.Sprintf("\n%s", ControlInputNameByCapability(howler.CapMouseAxis)),
  fmt.Sprintf("\n%s", ControlFunctionNameByCapability(howler.CapMouseAxis)),
)

func setMouseAxis(c *cli.Context) error {
  input := InputStruct{
    Command:   "set-mouse-axis",

    Name:      c.String("input"),
    Type:      c.String("axis"),
    //Value:     c.String(""),
    //Modifier:  c.String(""),
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
