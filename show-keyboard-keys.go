package main

import (
  "fmt"

  "gopkg.in/urfave/cli.v2"

  //howler "github.com/sferris/howler-controller"
)

func showKeyboardKeys(c *cli.Context) error {
  fmt.Println("Valid keyboard keys:\n")
  fmt.Println( KeyNames() );
  return nil
}

