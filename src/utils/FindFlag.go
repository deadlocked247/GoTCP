package utils

import (
  "bufio"
  "fmt"
  "net"
)

/**
 * Repeats until flag value is found from server
 * @param {net.Conn} conn - connection object
 * @param {bufio.Reader} reader - buffer that is coming from server
 * @return {string} flag - secret flag from server
 */
func FindFlag(conn net.Conn, reader bufio.Reader)(flag string){
  for {
    // Read in from string buffer from TCP connection
    response, err := reader.ReadString('\n')
    HandleError(err)

    // Check if API returns valid format provided by assignment
    HandleError(CheckFormat(response))

    // Parse API data to get flag and answer to question
    answer, flag, err := ParseData(response)
    HandleError(err)

    // Send response to server
    fmt.Fprintf(conn, "cs3700fall2016 %d\n", answer)
    if flag != "" {
      return flag
    }
  }
  return ""
}
