package main

import (
  "fmt"

  "gopkg.in/urfave/cli.v2"

  //howler "github.com/sferris/howler-controller"
)

func showKeyboardModifiers(c *cli.Context) error {
  fmt.Println("Valid keyboard modifier names:\n")
  fmt.Println( ModifierNames() );
  return nil
}

