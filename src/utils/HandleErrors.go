package utils

/**
 * Handles errors that are returned
 * @param {error} err - error to display and throw
 */
func HandleError(err error)() {
  if err != nil {
    panic(err)
  }
}
