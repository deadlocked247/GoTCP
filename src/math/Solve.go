package math

import (
	"errors"
	"strconv"
	)

/**
 * Does mathmatical caluclation, does not take negative values or values
 * greater than 1000
 * @param {int} a - first number to calculate
 * @param {string} operator - mathmatical operator to use, only accpet * / + -
 * @param {int} b - second number to calculate
 */
func Solve(a int, operator string, b int) (ans int, err error) {
	ans = 0
	if (a > 1000 || a < 0) {
		err = errors.New("First value is out of range" + strconv.Itoa(a))
		return ans, err
	}
	if (b > 1000 || b < 0) {
		err = errors.New("Second value is out of range" + strconv.Itoa(b))
		return ans, err
	}
	switch operator {
  	case "+":
  		ans = a + b
  	case "-":
  		ans = a - b
  	case "/":
  		ans = a / b
  	case "*":
  		ans = a * b
		default:
			err = errors.New("Unidentified operator" + operator)
  }
	return ans, err
}
