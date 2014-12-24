package main

import (
  "fmt"
  "net"
  "net/http"
  "os"
  "time"
)

func main() {
  tcpDaytime()
}

// Checks the nil status of an error, if it is non-nil
// it prints the error and quits the program
func errChk(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s: %s", time.Now().String(), err.Error())
    os.Exit(1)
  }
}

// Creates a daytime server that communicates using raw tcp packets
func tcpDaytime() {
  srv, err := net.Listen("tcp", ":2400")
  errChk(err)

  for {
    conn, err := srv.Accept()
    if err != nil { fmt.Fprint(os.Stderr, err) }

    go handleConn(conn)
  }
}

// Handles the raw tcp connections
func handleConn(conn net.Conn) {
  defer conn.Close()
  daytime := time.Now().String()
  conn.Write([]byte(daytime))
}

// Creates an http server that responds with the daytime string
func httpDaytime() {
  http.HandleFunc("/", httpHandler)
  http.ListenAndServe(":2400", nil)
}

// Handles the http connections
func httpHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, time.Now().String())
}
