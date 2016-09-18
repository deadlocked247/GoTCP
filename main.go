/**
 * Burak Aslan
 * Project 1
 * CS 3700
 */

package main

import (
  "fmt"
  "flag"
  "bufio"
  "reflect"
)

import (
  "./src/utils"
)

func main() {
    // Parse flags
    ssl := flag.Bool("s", false, "use ssl connection")
    port := flag.Int("p", 27993, "port to connect to")
    flag.Parse()

    if *ssl && *port == 27993 {
      *port = 27994
    }

    // Parse arguments from command line
    host, neuId, err := utils.ParseArgs()
    utils.HandleError(err)

    // Connect to TCP server
    conn, err := utils.Connect(host, *port, *ssl)
    utils.HandleError(err)

    // Send the first hello message
    fmt.Fprintf(conn, "cs3700fall2016 HELLO %s\n", neuId)

    // Start reading data
    reader := bufio.NewReader(conn)
    reflect.TypeOf(reader)
    fmt.Println(utils.FindFlag(conn, *reader))
}
