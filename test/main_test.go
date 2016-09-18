package main

import (
  "testing"
  "os"
)

import (
  "../src/utils"
  "../src/math"
)

func TestParseArgs(t *testing.T) {
  os.Args = []string{"run", "host", "neuId"}
  host, neuId, err := utils.ParseArgs()
  if err != nil { t.Error("Parse args failed on valid input") }
  if host != "host" || neuId != "neuId" { t.Error("Host domain or NeuID not set correctly") }
}

func TestParseData(t *testing.T) {
  ans, solution, err := utils.ParseData("class test test123")
  if solution != "test123" { t.Error("Solution was not returned correctly") }
  if ans != 0 { t.Error("Answer was set when solution was returned") }
  if err != nil { t.Error("Error'd on valid input") }
}

func TestSolve(t *testing.T) {
  ans, err := math.Solve(1, "+", 2)
  if ans != 3 { t.Error("Addition not correct") }

  ans, err = math.Solve(3, "-", 2)
  if ans != 1 { t.Error("Subtraction not correct") }

  ans, err = math.Solve(2, "*", 2)
  if ans != 4 { t.Error("Multiplication not correct") }

  ans, err = math.Solve(4, "/", 2)
  if ans != 2 { t.Error("Division not correct") }

  ans, err = math.Solve(1000000, "/", 1)
  if err == nil { t.Error("Expected out of range error")}

  ans, err = math.Solve(1, "Wtf?", 1)
  if err == nil { t.Error("Expected wrong operator error")}
}

func TestIsInArray(t *testing.T) {
  arr := []string{"test", "test123", "test456"}
  if !utils.IsInArray("test", arr) { t.Error("String was not found in array")}
  if utils.IsInArray("wrong", arr) { t.Error("Wrong string was found in array")}
}

func TestCheckFormat(t *testing.T) {
  err := utils.CheckFormat("nope")
  if err == nil { t.Error("No error thrown on invalid API data")}
  err = utils.CheckFormat("class BYE 123")
  if err != nil { t.Error("Error thrown on valid API data")}
}

func TestConnect(t *testing.T) {
  conn, err := utils.Connect("!3!@#!@$!#$#@%@#$", 123, false)
  if err == nil { t.Error("Connected to domain that cannot be connected to")}
  if conn != nil { t.Error("Failed to create conenction object")}
}
