package main

import (
  "os"
  "fmt"
  _"log"
  "flag"
  "io/ioutil"
  _"strings"
  _"strconv"

  _"time"

  "gopkg.in/yaml.v2"


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

  var data YAML
  source, err := ioutil.ReadFile(*FileName)

  if err != nil {
      panic(err)
  }
  err = yaml.Unmarshal(source, &data)
  if err != nil {
      panic(err)
  }
  fmt.Printf("%T\n", data.Leds[0])

  return nil
}

