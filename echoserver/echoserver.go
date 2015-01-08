package main

import (
  "net"
  "os"
  "fmt"
  "time"
)

func main() {
  laddr, err := net.ResolveTCPAddr("tcp", ":1221")
  errChk(err)

  srv, err := net.ListenTCP("tcp", laddr)
  errChk(err)

  for {
    conn, err := srv.Accept()
    if err != nil { continue }

    go handleConn(conn)
  }
}

func errChk(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s: %s", time.Now().String(), err.Error())
    os.Exit(1)
  }
}

func handleConn(conn net.Conn) {
  defer conn.Close()

  var buf [1024]byte
  for {
    rlen, err := conn.Read(buf[:])
    if err != nil { return }

    conn.Write(buf[:rlen])
  }
}
