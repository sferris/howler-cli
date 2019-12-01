package main

import (
  "fmt"

  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)

/*
  Name: "input",
  --
  Name: "modifier",
  --
  Name: "key",
*/

var txtKeyboardButton = `
This command will set a controller input to emit a keyboard key

- Inputs: %s

- Keys: %s

- Modifiers: %s
`

var usageKeyboardButton = "Set input to emit a keyboard key"
var descKeyboardButton = fmt.Sprintf(
  txtKeyboardButton,
  fmt.Sprintf("\n%s", ControlInputNameByCapability(howler.CapKeyboardButton)),
  fmt.Sprintf("\n%s", KeyNames()),
  fmt.Sprintf("\n%s", ModifierNames()),
)

func setKeyboardButton(c *cli.Context) error {
  input := InputStruct{
    Command:   "set-keyboard-button",

    Name:      c.String("input"),
    //Type:      c.String(""),
    Value:     c.String("key"),
    Modifier:  c.String("modifier"),
  }

  return input.Process()
}

// SetKeyboardButton(control ControlInput, key KeyCodes, modifier KeyModifiers)
func (input *InputStruct) SetKeyboardButton() error {
  var err error

  control, err := howler.StringToControl(input.Name)
  if err != nil {
    return err
  }

  key := howler.ToKey(input.Value)
  if key == -1 {
    return fmt.Errorf(
      "Invalid keyboard code for input: '%s': ",
      input.Value,
    )
  }

  modifier := howler.ToModifier(input.Modifier)
  if modifier == -1 {
    return fmt.Errorf(
      "Invalid keyboard modifier for input: '%s': ",
      input.Modifier,
    )
  }

  if controller == nil { 
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return err
    }
  }   

  _, err = controller.SetKeyboardButton(control, key, modifier)
  return err
}
