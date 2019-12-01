package main

import (
  "fmt"

  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)

/*
  Name: "input",
  --
  Name: "joystick"
  Default: "joystick1",
*/

func setJoystickButton(c *cli.Context) error {
  input := InputStruct{
    Name:      c.String(""),
    Type:      c.String(""),
    Modifier:  c.String(""),
    Value:     c.String(""),
  }

  return input.Process()
}

// SetJoystickButton(control ControlInput, joystick ControlFunction, button JoystickButtons)
func (input *InputStruct) SetJoystickButton() error {
  var err error

  control, err := howler.StringToControl(input.Name)
  if err != nil {
    return err
  }

  function, err := howler.StringToControlFunction(input.Type)
  if err != nil {
    return err
  }

  button := howler.ToJoystickButton(input.Value)
  if button == -1 {
    return fmt.Errorf(
      "Invalid joystick button number for input: '%s'",
      input.Value,
    )
  }

  if controller == nil { 
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return err
    }
  }   
    
  _, err = controller.SetJoystickButton(control, function, button)
  return err
}
