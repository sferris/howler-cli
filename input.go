package main

import (
  "os"
  "fmt"
  "strings"
  "flag"

  howler "github.com/sferris/howler-controller"
)

type Input struct {
  Button    string
  Mode      string
  Modifier  string
  Value     string
}

var setInputCMD *flag.FlagSet

func init() {
  setInputCMD = flag.NewFlagSet("set-input", flag.ExitOnError)
}

func parseInput() error {
  InputButton   := setInputCMD.String("input", "", "Button to set the Input properties (Required)")
  InputMode     := setInputCMD.String("mode", "", "The Input mode [Joystick1 or 2||Keyboard||Mouse] (Required)")
  InputModifier := setInputCMD.String("modifier", "", "The Input key modifier")
  InputValue    := setInputCMD.String("value", "", "The Input value (Required)")

  setInputCMD.Parse(os.Args[2:])

  if setInputCMD.Parsed() {
    input := Input{
      Button:    *InputButton,
      Mode:      *InputMode,
      Modifier:  *InputModifier,
      Value:     *InputValue,
    }

    input.Set();
  }

  return nil
}

func (input *Input) String() string {
  switch strings.ToLower(input.Mode) {
    case "joystick1": fallthrough
    case "joystick2":
      return fmt.Sprintf(
        "Button: %d, Mode: %d, Modifier: %d, Value: %d",
          howler.Input(input.Button),
          howler.Mode(input.Mode),
          howler.Modifier(input.Modifier),
          input.Value,
      )
    case "keyboard":
      return fmt.Sprintf(
        "Button: %d, Mode: %d, Modifier: %d, Value: %d",
          howler.Input(input.Button),
          howler.Mode(input.Mode),
          howler.Modifier(input.Modifier),
          howler.Key(input.Value),
      )
    case "mouse":
      return fmt.Sprintf(
        "Button: %d, Mode: %d, Modifier: %d, Value: %d",
          howler.Input(input.Button),
          howler.Mode(input.Mode),
          howler.Modifier(input.Modifier),
          howler.MouseButton(input.Value),
      )
  }
  return ""
}

func (input *Input) Set() error {
  fmt.Println(input)
  return nil
}
