package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    port := ":8081"
    buffsize := 4096
    exit_message := "q"
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
