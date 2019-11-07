package main

import (
  "os"
  "fmt"
  _"log"
  "flag"
  _"strings"
  _"strconv"

  _"time"

  //howler "github.com/sferris/howler-controller"
)

type YAML struct {
  Game    string
  Leds    []Led
  Inputs  []Input
}

var setReadFileCMD *flag.FlagSet

func init() {
  setReadFileCMD = flag.NewFlagSet("read-file", flag.ExitOnError)
}

func parseYAML() error {
  FileName := setReadFileCMD.String("file", "", "File to read instructions from (Required)")

  setReadFileCMD.Parse(os.Args[2:])

  if setReadFileCMD.Parsed() {
    fmt.Printf("Filename: %s\n", *FileName)
  }

  return nil
}

