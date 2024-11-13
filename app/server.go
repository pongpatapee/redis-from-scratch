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

		if value.typ != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}

		if len(value.array) == 0 {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)
		fmt.Println("Hanlding command", strings.ToUpper(command))

		args := value.array[1:]
		writer := NewWriter(conn)

		handler, ok := CommandHanlders[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(Value{typ: "string", str: ""})
		}

		result := handler(args)
		writer.Write(result)
	}
}
