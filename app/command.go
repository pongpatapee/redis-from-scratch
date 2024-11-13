package main

import (
	"strconv"
	"strings"
	"time"
)

var CommandHanlders = map[string]func([]Value) Value{
	"PING": PingHanlder,
	// "ECHO": EchoHandler,
	"SET": SetHandler,
	"GET": GetHandler,
}

func PingHanlder(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}

	// in redis if arg is provided to PING it echos
	return Value{typ: "string", str: args[0].bulk}
}

func EchoHandler(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'echo' command"}
	}

	return Value{typ: "string", str: args[0].bulk}
}

func SetHandler(args []Value) Value {
	if len(args) < 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'set' command"}
	}

	// TODO: Handle arguments better
	var opt string
	if len(args) >= 3 {
		opt = args[2].bulk
	}

	key := args[0].bulk
	value := args[1].bulk

	StringDB.Set(key, value)

	if strings.ToUpper(opt) == "PX" {
		if len(args) < 4 {
			return Value{typ: "error", str: "ERR PX missing arg"}
		}

		length, _ := strconv.Atoi(args[3].bulk)

		timer := time.NewTimer(time.Duration(length) * time.Millisecond)

		go func() {
			<-timer.C
			StringDB.Del(key)
		}()

	}

	return Value{typ: "string", str: "OK"}
}

func GetHandler(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].bulk

	val, exist := StringDB.Get(key)
	if !exist {
		return Value{typ: "null"}
	}

	// Marshaller automatically appends the length
	return Value{typ: "bulk", bulk: val}
}
