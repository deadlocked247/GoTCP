package utils

import (
  "os"
  "errors"
)

/**
 * Parses arguments from command line, while ignoring flags -p and -s
 * @return {string} host - domain to conencto to
 * @return {string} neuId - neu id to submit on
 * @return {error} err - error if length of arguments is incorrect
 */
func ParseArgs()(host string, neuId string, err error) {
  args := os.Args[1:]
  var fl string
  if len(args) < 2 {
    err = errors.New("Not enough arguments provided")
  }

  for i := 0; i < len(args); i++ {
    if len(args[i]) > 2 {
      fl = args[i][0:2]
    }
    if fl != "" && fl != "-p" && fl != "-s" {
      if (host != "") {
        neuId = args[i]
      } else {
        host = args[i]
      }
      fl = ""
    }
  }
  return host, neuId, err
}
