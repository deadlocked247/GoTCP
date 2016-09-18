package utils

import (
  "strings"
  "errors"
)

/**
 * Checks the format is valid from API
 * @param {[]string} data - split array of strings
 * @return {error} err - if it is invalid
 */
func CheckFormat(data string)(err error) {
  response := strings.Split(strings.TrimSpace(data), " ")
  correctFormats := []string{"HELLO", "STATUS", "SOLUTION", "BYE"}
  if len(response) < 3 {
    return errors.New("Server returned malformed data " + data)
  }
  if !IsInArray(response[1], correctFormats) {
    return errors.New("Server returned incorrect format of status " + response[1])
  }
  return nil
}
