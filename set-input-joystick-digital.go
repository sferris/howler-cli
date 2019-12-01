package main

import (
  "fmt"
  "strconv"

  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)


/*
  Name: "input",
  --
  Name: "function",
  --
  Name: "value",
*/

func setJoystickDigital(c *cli.Context) error {
  input := InputStruct{
    Command:   "set-joystick-digital",

    Name:      c.String("input"),
    Type:      c.String("function"),
    Value:     c.String("value"),
    //Modifier:  c.String(""),
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

  value, err := strconv.ParseInt(input.Value, 10, 8)
  if err != nil {
    return err
  }

  if controller == nil { 
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return fmt.Errorf("Unable to parse value: %s\n", err.Error())
    }
  }   
    
  _, err = controller.SetJoystickDigital(control, function, int8(value))
  return err
}
