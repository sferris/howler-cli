package main

import (
  "os"
  "log"

  "gopkg.in/urfave/cli.v2"
  howler "github.com/sferris/howler-controller"
)

var (
  device int
)

var app = &cli.App{
    Name: "howler-cli",
    Usage: "A command line utility for configuring a Howler Arcade controller",
    Version: "0.0.1",

    Flags: []cli.Flag{
      &cli.IntFlag{
        Name: "device",
        Aliases: []string{"d"},
        Value: 0,
        Usage: "The Howler Controller to configure (0..n)",
        Destination: &device,
      },
    },

    Commands: []*cli.Command{
      {
        Name:        "set-led",
        Usage:       "Change the color of one of the Button LEDs",
        Description: "This command is used to alter the color of a button/LED",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "button",
            Usage: "The button/led to change",
          },
          &cli.StringFlag{
            Name: "color",
            Usage: "The color to set the button/led to",
          },
          &cli.StringFlag{
            Name: "scope",
            Usage: "The scope to change the color (current, or default)",
          },
        },

        Action: func(c *cli.Context) error {
          led := LEDStruct{
            Name:   c.String("button"),
            Color:  c.String("color"),
            Scope:  c.String("scope"),
          }
          return led.Process()
        },
      },

      {
        Name:        "set-input",
        Usage:       "Alter the behavior of an input on the controller",
        Description: "This command is used to alter the behavior of an input on the controller",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "input",
            Usage: "The input to alter, eg: joy1up, joy1down, button1, etc",
          },
          &cli.StringFlag{
            Name: "mode",
            Usage: "The context used by the button when emitting a value (joystick1 or 2, keyboard, or mouse)",
          },
          &cli.StringFlag{
            Name: "modifier",
            Usage: "In keyboard mode, the modifier to use in addition to the value. ([left|right] control, shift, alt, or ui)",
          },
          &cli.StringFlag{
            Name: "value",
            Usage: "What value the button emits when pressed (context dependent)",
          },
        },

        Action: func(c *cli.Context) error {
          input := InputStruct{
            Name:      c.String("input"),
            Mode:      c.String("mode"),
            Modifier:  c.String("modifier"),
            Value:     c.String("value"),
          }
          return input.Process()
        },
      },

      {
        Name:        "from-file",
        Usage:       "Read settings from a yaml file",
        Aliases:     []string{"read"},
        Description: "This command allows you to change many settings that are represented in a yaml file",
        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "path",
            Aliases: []string{"file"},
            Usage: "The fullpath to the yaml file containing the settings to be applied",
          },
        },

        Action: func(c *cli.Context) error {
          file := FileStruct{
            Path: c.String("path"),
          }
          return file.Process()
        },

      },
    },
}

var controller *howler.HowlerDevice

func main() {
  var err error
  controller, err = howler.OpenDevice(device)
  if err != nil {
    log.Fatal(err.Error())
  }

  defer func() { controller.Close() }()

  err = app.Run(os.Args)
  if err != nil {
    log.Fatal(err.Error())
  }
}
