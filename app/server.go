package main

import (
	"fmt"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var (
	_ = net.Listen
	_ = os.Exit
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
			// os.Exit(1)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	for {

		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			return
		}

		// fmt.Println("Received Data:", string(buf[:n]))

		// // return back the same data
		// conn.Write(buf[:n])
		conn.Write([]byte("+PONG\r\n"))
	}
}
