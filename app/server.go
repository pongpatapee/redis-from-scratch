package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
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

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Starting connection!")

	for {

		buf := make([]byte, 1024)

		length, err := conn.Read(buf)
		if err == io.EOF {
			return
		} else if err != nil {
			fmt.Println("Error while trying to read from connection", err.Error())
			return
		}

		commandData := string(buf[:length])
		fmt.Println("Command Data")
		fmt.Println(commandData)
		args, _, err := parseCommand(commandData)
		if err != nil {
			fmt.Println("Error while trying to parse command")
			return
		}

		command := args[0]

		switch command {
		case "ECHO":
			fmt.Println("processing echo command")
			if len(args) != 2 {
				conn.Write([]byte("-ERR wrong number of arguments for 'echo' command"))
				return
			}

			reply := fmt.Sprintf("+%v\r\n", args[1])

			conn.Write([]byte(reply))
		case "PING":
			conn.Write([]byte("+PONG\r\n"))
		default:
			conn.Write([]byte("-ERR unknown command\r\n"))
			// return
		}

		// fmt.Println("Parsed args")
		// fmt.Println(args, len(args))
		// fmt.Println("Parsed length")
		// fmt.Println(argLengths, len(argLengths))

	}
}

func parseCommand(commandData string) ([]string, []int, error) {
	metadata := strings.Split(commandData, "\r\n")

	numArgs, err := strconv.Atoi(metadata[0][1:])
	if err != nil {
		fmt.Println("Error extracting num args")
		return nil, nil, err
	}

	args := make([]string, numArgs)
	argLengths := make([]int, numArgs)
	idx := 1
	for i := range numArgs {
		argLength, _ := strconv.Atoi(metadata[idx][1:])
		argLengths[i] = argLength
		args[i] = metadata[idx+1]
		idx += 2
	}

	// make command case-insensitive
	args[0] = strings.ToUpper(args[0])

	return args, argLengths, nil
}
