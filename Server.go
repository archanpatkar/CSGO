package main

import (
    "fmt"
    "net"
    "os"
    "strconv"
)

func main() {
    var port string
    var buffsize int
    var exit_message string
    if os.Args[1] != "" {
      port = ":" + os.Args[1]
    }
    if os.Args[2] != "" {
      number, _ := strconv.Atoi(os.Args[2])
      buffsize = number
    }
    if os.Args[3] != "" {
      exit = number
    }
    fmt.Println("Server Starting on PORT " + port)
    address, err := net.ResolveTCPAddr("tcp4", port)
    checkError(err)
    ss, err := net.ListenTCP("tcp", address)
    checkError(err)
    for {
        conn, err := ss.AcceptTCP()
        if err != nil {
            continue
        }
        go handler(conn,buffsize,exit_message)
    }
}

func handler(conn * net.TCPConn, buffsize int, exit_message string) {
  // Boiler Plate Code
  data := make([]byte, buffsize)
  _ ,err := conn.Read(data)
  message := string(data[:buffsize])
  if err != nil {
      fmt.Fprintf(os.Stderr, "Client Error: %s", err.Error())
      goto End
  }
  fmt.Println(message)
  for message != exit_message {
    data = make([]byte, buffsize)
    _ ,err := conn.Read(data)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Client Error: %s", err.Error())
        goto End
    }
    message = string(data[:buffsize])
    fmt.Println(message)
  }
  End: conn.Close();
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
        os.Exit(1)
    }
}
