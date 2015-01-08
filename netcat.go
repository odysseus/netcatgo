package main

import (
  "net"
  "fmt"
  "bufio"
  "log"
  "os"
)

func main() {
}

func Listen(conn net.Conn) {
  var buf[1024] byte
  for {
    rlen, err := conn.Read(buf[:])
    if err != nil { fmt.Println(err) }

    s := string(buf[:rlen])
    fmt.Print(s)
  }
}

func Transmit(conn net.Conn) {
  reader := bufio.NewReader(os.Stdin)
  for {
    input, err := reader.ReadBytes('\n')
    if err != nil { fmt.Println(err) }

    conn.Write(input)
  }
}

func Communicate(conn net.Conn) {
  defer conn.Close()
  go Listen(conn)
  Transmit(conn)
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

func TCPConnect(laddr, raddr string) (*net.TCPConn, error) {
  // Resolve the local address
  lad, err := net.ResolveTCPAddr("tcp", laddr)
  if err != nil { return nil, err }

  // Resolve the remote address
  rad, err := net.ResolveTCPAddr("tcp", raddr)
  if err != nil { return nil, err }

  // Make the connection
  conn, err := net.DialTCP("tcp", lad, rad)
  if err != nil { return nil, err }

  return conn, nil
}

func TCPListener(laddr string) {
  l, err := net.Listen("tcp", laddr)
  if err != nil { log.Fatal(err) }

  defer l.Close()
  for {
    // Wait for a connection
    conn, err := l.Accept()
    if err != nil { log.Fatal(err) }

    // Handle connections in a new goroutine
    go func(c net.Conn) {
      for {
        reader := bufio.NewReader(c)

        s, err := reader.ReadBytes('\n')
        if err != nil { fmt.Println(err) }
        c.Write(s)
      }

      c.Close()
    }(conn)
  }
}
