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


var txtJoystickDigital = `
This command will set a controller input to emit digital joystick movements

Inputs: %s

Functions: %s

Value: Any number between -127 and 127
`

var usageJoystickDigital = "Set input to emit digital joystick movements"
var descJoystickDigital = fmt.Sprintf(
  txtJoystickDigital,
  fmt.Sprintf("\n%s", ControlInputNameByCapability(howler.CapJoystickDigital)),
  fmt.Sprintf("\n%s", ControlFunctionNameByCapability(howler.CapJoystickDigital)),
)

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
