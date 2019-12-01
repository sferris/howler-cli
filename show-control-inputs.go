package main

import (
  "fmt"

  "gopkg.in/urfave/cli.v2"

  //howler "github.com/sferris/howler-controller"
)

func showControlInputs(c *cli.Context) error {
  fmt.Println("Valid Control inputs:\n")
  fmt.Println( ControlInputNames() );

  return nil
}

