package utils

/**
 * Check if a string is in an array of strings
 * @param {string} str - string to find
 * @param {[]string} arr - array of strings
 * @return bool
 */
func IsInArray(str string, arr []string) bool {
  for _, x := range arr {
    if x == str {
      return true
    }
  }
  return false
}
