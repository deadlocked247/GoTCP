package utils

import (
  "strings"
  "strconv"
  "errors"
)

import (
  "../math"
)

/**
 * Parses data from TCP connection
 * @param {string} data - The string TCP server returns
 * @return {int} ans - integer value from solving problem returned
 * @return {string} flag - secret flag value
 * @return {error} err - error if parsing is invalid
 */
func ParseData(data string) (ans int, flag string, err error) {
	splitArr := strings.Split(strings.TrimSpace(data), " ")
  ans = 0
  flag = ""
	if len(splitArr) >= 5 {
		a, _ := strconv.Atoi(splitArr[2])
		b, _ := strconv.Atoi(splitArr[4])
		ans, err = math.Solve(a, splitArr[3], b)
    HandleError(err)
	} else if len(splitArr) > 1 {
		flag = splitArr[2]
	} else {
    err = errors.New("Not enough arguments")
  }
	return ans, flag, err
}
