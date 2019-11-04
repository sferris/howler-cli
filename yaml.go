package main

import (
  _"io/ioutil"
  _"gopkg.in/yaml.v2"
)

type Input struct {
  Button    string
  Mode      string
  Modifier  string
  Value     string
}

type Led struct {
  Button    string
  Mode      string
  RGB       []int
}
  
type Work struct {
  Game    string
  Leds    []Led
  Inputs  []Input
}

/*
func main() {
  var work Work

  source, err := ioutil.ReadFile("mario.yaml")
  if err != nil {
    panic(err)
  }  

  err = yaml.Unmarshal(source, &work)
  if err != nil {
    panic(err)
  }

  fmt.Printf("Value: %#v\n", work)

  os.Exit(0)
}
*/
