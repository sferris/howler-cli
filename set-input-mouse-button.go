package main

import (
  "fmt"

  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)

/*
  Name: "input",
  --
  Name: "button",
*/

func setMouseButton(c *cli.Context) error {
  input := InputStruct{
    Name:      c.String(""),
    Type:      c.String(""),
    Modifier:  c.String(""),
    Value:     c.String(""),
  }

  return input.Process()
}

// SetMouseButton(control ControlInput, button MouseButtons)
func (input *InputStruct) SetMouseButton() error {
  var err error

  control, err := howler.StringToControl(input.Name)
  if err != nil {
    return err
  }

  button := howler.ToMouseButton(input.Value)
  if button == -1 {
    return fmt.Errorf(
      "Invalid mouse button number for input: '%s': ",
      input.Value,
    )
  }

  if controller == nil { 
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return err
    }
  }   
    
  _, err = controller.SetMouseButton(control, button)
  return err
}
