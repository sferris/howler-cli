package main

import (
  "os"
  "log"
  _"fmt"

  _"sort"

  "syscall"
  "golang.org/x/sys/unix"

  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)

var (
  device int
)

var columns = 70

func init() {
  ws, err := unix.IoctlGetWinsize(syscall.Stdout, unix.TIOCGWINSZ)
  if err == nil {
    columns = int(ws.Col)
  }
}

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
      // \
      //  > LED Actions
      // /

      // Set LED color
      {
        Name:        "set-led",
        Usage:       "Change the color of one of the Button LEDs",
        Description: "This command is used to alter the color of a button/LED",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "input",
            Usage: "The button/led to change",
          },
          &cli.StringFlag{
            Name: "color",
            Usage: "The color to set the button/led to",
          },
        },

        Action: setLED,
      },

      // Set LED default color
      {
        Name:        "set-led-default",
        Usage:       "Change the color of one of the Button LEDs",
        Description: "This command is used to alter the color of a button/LED",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "input",
            Usage: "The button/led to change",
          },
          &cli.StringFlag{
            Name: "color",
            Usage: "The color to set the button/led to",
          },
        },

        Action: setDefaultLED,
      },

      // \
      //  > Input Actions
      // /

      // Set input to emit analog joystick movement
      {
        Name:        "set-joystick-analog",
        Aliases:     []string{"setja"},
        Usage:       "Set input to emit analog joystick movements",
        Description: "This command will set a controller input to emit analog joystick movements",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "input",
            Usage: "The input to alter, eg: joy1up, joy1down, button1, etc",
          },
          &cli.StringFlag{
            Name: "function",
            Usage: "Which mouse axis to emit",
          },
        },

        Action: setJoystickAnalog,
      },

      // Set input to emit joystick button codes
      {
        Name:        "set-joystick-button",
        Aliases:     []string{"setjb"},
        Usage:       "Set input to emit a joystick button",
        Description: "This command will set a controller input to emit a joystick button press",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "input",
            Usage: "The input to alter, eg: joy1up, joy1down, button1, etc",
          },
          &cli.StringFlag{
            Name: "joystick",
            Value: "joystick1",
            Usage: "The joystick number to emit, eg: 1 or 2",
          },
          &cli.StringFlag{
            Name: "button",
            Usage: "Which joystick button to emit",
          },
        },

        Action: setJoystickButton,
      },

      // Set input to emit digital joystick movement
      {
        Name:        "set-joystick-digital",
        Aliases:     []string{"setjd"},
        Usage:       "Set input to emit digital joystick movements",
        Description: "This command will set a controller input to emit digital joystick movements",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "input",
            Usage: "The input to alter, eg: joy1up, joy1down, button1, etc",
          },
          &cli.StringFlag{
            Name: "function",
            Usage: "Which mouse axis to emit",
          },
          &cli.StringFlag{
            Name: "value",
            Usage: "Value for axis (-128 .. 128)",
          },
        },

        Action: setJoystickDigital,
      },

      // Set input to emit keyboard key-codes
      {
        Name:        "set-keyboard-button",
        Aliases:     []string{"setkb"},
        Usage:       "Set input to emit a keyboard key",
        Description: "This command will set a controller input to emit a keyboard key",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "input",
            Usage: "The input to alter, eg: joy1up, joy1down, button1, etc",
          },
          &cli.StringFlag{
            Name: "modifier",
            Usage: "In a keyboard context, the modifier to use in addition to the value",
          },
          &cli.StringFlag{
            Name: "key",
            Usage: "What value the button emits when pressed (context dependent)",
          },
        },

        Action: setKeyboardButton,
      },

      // Set input to emit mouse button codes
      {
        Name:        "set-mouse-axis",
        Aliases:     []string{"setma"},
        Usage:       "Set input to emit mouse movements",
        Description: "This command will set a controller input to emit a mouse movements",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "input",
            Usage: "The input to alter, eg: joy1up, joy1down, button1, etc",
          },
          &cli.StringFlag{
            Name: "axis",
            Usage: "Which mouse axis to emit",
          },
        },

        Action: setMouseAxis,
      },

      // Set input to emit mouse button codes
      {
        Name:        "set-mouse-button",
        Aliases:     []string{"setmb"},
        Usage:       "Set input to emit a mouse button",
        Description: "This command will set a controller input to emit a mouse button press",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "input",
            Usage: "The input to alter, eg: joy1up, joy1down, button1, etc",
          },
          &cli.StringFlag{
            Name: "button",
            Usage: "Which mouse button to emit",
          },
        },

        Action: setMouseButton,
      },
      // \
      //  > Miscelaneous
      // /

      // Reset to defaults
      {
        Name:        "reset-defaults",
        Aliases:     []string{"defaults"},
        Usage:       "Reset controller to defaults",
        Description: "This command will reset the controller to defaults",

        Action: resetDefaults,
      },
      // \
      //  > Process many inputs defined in a file
      // /
      {
        Name:        "read-file",
        Usage:       "Read settings from a yaml file",
        Aliases:     []string{"read","file"},
        Description: "This command allows you to change many settings that are represented in a yaml file",
        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "path",
            Aliases: []string{"file"},
            Usage: "The fullpath to the yaml file containing the settings to be applied",
          },
          &cli.IntFlag{
            Name: "sleep",
            Usage: "Number of milliseconds to sleep between LED color changes",
          },
        },

        Action: readFile,
      },

      // Show controller LED settings
      {
        Name:        "show-led-settings",
        Aliases: []string{"led-settings"},
        Usage:       "Display the controller LED strtings",
        Description: "Display the controller LED strtings",

        Action: showLEDSettings,
      },

      // Get controller Input settings
      {
        Name:        "show-control-settings",
        Aliases: []string{"control-settings"},
        Usage:       "Display the controller LED strtings",
        Description: "Display the controller LED strtings",

        Action: showControlSettings,
      },

      // \
      //  > Show possible keyboard key values
      // /
      {
        Name:        "show-keyboard-keys",
        Usage:       "Show the valid keyboard keys",
        Aliases:     []string{"show-keys"},
        Description: "This shows the list of valid keyboard key names",

        Action: showKeyboardKeys,
      },

      // \
      //  > Show possible keyboard modifier values
      // /
      {
        Name:        "show-keyboard-modifiers",
        Usage:       "Show the valid keyboard modifier names",
        Aliases:     []string{"show-modifiers"},
        Description: "This shows the list of valid keyboard modifier names",

        Action: showKeyboardModifiers,
      },

      // \
      //  > Show possible mouse button values
      // /
      {
        Name:        "show-mouse-buttons",
        Usage:       "Show the valid mouse buttons",
        Aliases:     []string{"show-mouse"},
        Description: "This shows the list of valid mouse buttons",

        Action: showMouseButtons,
      },

      // \
      //  > Show possible joystick button values
      // /
      {
        Name:        "show-joystick-buttons",
        Usage:       "Show the valid joystick buttons",
        Aliases:     []string{"show-joystick"},
        Description: "This shows the list of valid joystick buttons",

        Action: showJoystickButtons,
      },

      // \
      //  > Show possible inputs
      // /
      {
        Name:        "show-control-inputs",
        Usage:       "Show the valid control inputs (context dependent)",
        Aliases:     []string{"showci"},
        Description: "This will show the possible inputs for configuring inputs, or settings LEDs",

        Action: showControlInputs,
      },


      // \
      //  > Show firmware version
      // /
      {
        Name:        "show-firmware-revision",
        Usage:       "Show the firmware revision information",
        Aliases:     []string{"show-fw","firmware"},
        Description: "This shows the list of valid joystick buttons",

        Action: showFirmwareRevision,
      },
    },
}

var controller *howler.HowlerDevice

func main() {
  log.SetFlags(0)
  //log.SetOutput(ioutil.Discard)

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err.Error())
  }

  if controller != nil {
    controller.Close()
  }
}

