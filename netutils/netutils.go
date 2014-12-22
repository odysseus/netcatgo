/// Short examples of useful methods

package main

import (
  "fmt"
  "net"
)

func main() {
  s, _ := GetHTTPHead("google.com:80")
  fmt.Println(s)
}

// Prints the host IP address for a given URI
func PrintHost(name string) {
  addrs, err := net.LookupHost(name)
  if err != nil { fmt.Println(err); return }

  for _, s := range addrs {
    fmt.Println(s)
  }
}

// Prints the canonical DNS host for the given URI
func PrintCName(name string) {
  cname, err := net.LookupCNAME(name)
  if err != nil { fmt.Println(err); return }

  fmt.Println(cname)
}

// Opens a TCP connection to addr and sends a request for the HTTP head
func GetHTTPHead(addr string) (string, error){
  raddr, err := net.ResolveTCPAddr("tcp", addr)
  if err != nil { return "", err }

  conn, err := net.DialTCP("tcp", nil, raddr)
  if err != nil { return "", err }

  conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
  var buffer [1024]byte
  rlen, _ := conn.Read(buffer[:])

  return string(buffer[:rlen]), nil
}
