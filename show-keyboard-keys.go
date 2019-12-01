package main

import (
  "fmt"
  "sort"

  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)

func showKeyboardKeys(c *cli.Context) error {
  fmt.Println("Valid keyboard keys:\n")
  fmt.Println( KeyNames() );
  return nil
}

func KeyNames() string {
  var result string

  var keys []int
  for k := range howler.KeyNames {
    keys = append(keys, int(k))
  }

  sort.Ints(keys)

  w := 0
  for _, k := range keys {
    value := fmt.Sprintf("%s, ", howler.KeyNames[howler.KeyCodes(k)])

    result += value

    w += len(value)
    if w >= (columns-20) {
      w=0
      result += fmt.Sprintln()
    }

  }

  return result + fmt.Sprintln()
}
