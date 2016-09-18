package utils
import (
  "net"
  "strconv"
  "crypto/tls"
)

/**
 * Connect to TCP server
 * @param {string} host - domain to connect to
 * @param {string} port - port to use
 * @param {bool} ssl - use SSL to connect
 * @return {net.Conn} conn - connection to server
 * @return {error} err - error if connection failed
 */
func Connect(host string, port int, ssl bool)(conn net.Conn, err error) {
  if ssl {
    config := &tls.Config {
      InsecureSkipVerify: true,
    }
    return tls.Dial("tcp", host + ":" + strconv.Itoa(port), config)
  } else {
    return net.Dial("tcp", host + ":" + strconv.Itoa(port))
  }
  return nil, nil
}
