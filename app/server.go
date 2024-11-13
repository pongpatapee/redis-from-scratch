package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("Server listening on port 6379")

	for {
		conn, err := listener.Accept() // blocking call
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Starting connection!")

	for {

		resp := NewResp(conn)

		value, err := resp.Read()
		if err != nil {
			if err == io.EOF {
				return
			}

			fmt.Println("Error while trying to read from connection", err.Error())
			return
		}

		args := make([]string, len(value.array))
		for i, v := range value.array {
			args[i] = v.bulk
		}

		command := args[0]
		fmt.Println("Hanlding command", strings.ToUpper(command))

		handler, ok := CommandHanlders[command]
		if ok {
			handler(conn, args[1:])
		} else {
			conn.Write([]byte("-ERR unknown command\r\n"))
		}

	}
}
