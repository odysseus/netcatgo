package main

import (
  "net"
  "fmt"
  "bufio"
  "os"
)

func main() {
  conn, err := UDPConnect("localhost:5050", ":4040")
  if err != nil { fmt.Println(err); return }
  UDPCommunicate(conn)
}

func UDPConnect(locaddr, remaddr string) (*net.UDPConn, error) {
  // Resolving the local address
  laddr, err := net.ResolveUDPAddr("udp", locaddr)
  if err != nil { return nil, err }

  // Resolving the remote address
  raddr, err := net.ResolveUDPAddr("udp", remaddr)
  if err != nil { return nil, err }

  // Getting the connection
  conn, err := net.DialUDP("udp", laddr, raddr)
  if err != nil { return nil, err }

  return conn, nil
}

func UDPListen(conn *net.UDPConn) {
  var buf[1024] byte
  for {
    rlen, err := conn.Read(buf[:])
    if err != nil { fmt.Println(err) }

    s := string(buf[:rlen])
    fmt.Print(s)
  }
}

func UDPTransmit(conn *net.UDPConn) {
  reader := bufio.NewReader(os.Stdin)
  for {
    input, err := reader.ReadBytes('\n')
    if err != nil { fmt.Println(err) }
    conn.Write(input)
  }
}

func UDPCommunicate(conn *net.UDPConn) {
  go UDPListen(conn)
  UDPTransmit(conn)
}
