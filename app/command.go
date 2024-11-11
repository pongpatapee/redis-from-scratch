package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var CommandHanlders = map[string]func(net.Conn, []string){
	"PING": PingHanlder,
	"ECHO": EchoHandler,
	"SET":  SetHandler,
	"Get":  GetHandler,
}

func ParseCommand(commandData string) ([]string, []int, error) {
	metadata := strings.Split(commandData, "\r\n")

	if len(metadata) < 3 || metadata[0][0] != '*' {
		return nil, nil, errors.New("invalid RESP array format")
	}

	numArgs, err := strconv.Atoi(metadata[0][1:])
	if err != nil {
		fmt.Println("Error extracting num args")
		return nil, nil, err
	}

	args := make([]string, numArgs)
	argLengths := make([]int, numArgs)
	idx := 1
	for i := range numArgs {

		if len(metadata) == 0 || metadata[idx][0] != '$' {
			return nil, nil, errors.New("expected RESP bulk string")
		}

		argLength, err := strconv.Atoi(metadata[idx][1:])
		if err != nil {
			return nil, nil, err
		}

		argLengths[i] = argLength
		args[i] = metadata[idx+1]
		idx += 2
	}

	// make command case-insensitive
	args[0] = strings.ToUpper(args[0])

	return args, argLengths, nil
}

func PingHanlder(conn net.Conn, args []string) {
	conn.Write([]byte("+PONG\r\n"))
}

func EchoHandler(conn net.Conn, args []string) {
	if len(args) != 1 {
		conn.Write([]byte("-ERR wrong number of arguments for 'echo' command\r\n"))
		return
	}

	reply := fmt.Sprintf("+%v\r\n", args[0])

	conn.Write([]byte(reply))
}

func SetHandler(conn net.Conn, args []string) {
}

func GetHandler(conn net.Conn, args []string) {
}
