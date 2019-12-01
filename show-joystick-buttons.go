package main

import (
  "fmt"
  "sort"

  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)

func showJoystickButtons(c *cli.Context) error {
  fmt.Println("Valid joystick buttons:\n")
  fmt.Println( JoystickButtons() );
  return nil
}

func JoystickButtons() string {
  var result string

  var keys []int
  for k := range howler.JoystickButtonNames {
    keys = append(keys, int(k))
  }

  sort.Ints(keys)

  w := 0
  for _, k := range keys {
    value := fmt.Sprintf("%s, ", howler.JoystickButtonNames[howler.JoystickButtons(k)])

    result += value

    w += len(value)
    if w >= (columns-20) {
      w=0
      result += fmt.Sprintln()
    }

  }

  return result + fmt.Sprintln()
}
