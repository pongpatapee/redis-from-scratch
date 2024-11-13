package main

var CommandHanlders = map[string]func([]Value) Value{
	"PING": PingHanlder,
	// "ECHO": EchoHandler,
	// "SET":  SetHandler,
	// "GET":  GetHandler,
}

func PingHanlder(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}

	// in redis if arg is provided to PING it echos
	return Value{typ: "string", str: args[0].bulk}
}

// func EchoHandler(args []Value) Value {
// 	if len(args) != 1 {
// 		conn.Write([]byte("-ERR wrong number of arguments for 'echo' command\r\n"))
// 		return
// 	}

// 	reply := fmt.Sprintf("+%v\r\n", args[0])

// 	conn.Write([]byte(reply))
// }

// func SetHandler(args []Value) Value {
// 	if len(args) < 2 {
// 		conn.Write([]byte("-ERR syntax error\r\n"))
// 		return
// 	}

// 	var opt string
// 	if len(args) >= 3 {
// 		opt = args[2]
// 	}

// 	key := args[0]
// 	value := args[1]

// 	StringStore.Set(key, value)

// 	if strings.ToUpper(opt) == "PX" {
// 		if len(args) < 4 {
// 			conn.Write([]byte("-ERR PX missing arg\r\n"))
// 			return
// 		}

// 		length, _ := strconv.Atoi(args[3])

// 		timer := time.NewTimer(time.Duration(length) * time.Millisecond)

// 		go func() {
// 			<-timer.C
// 			StringStore.Del(key)
// 		}()

// 	}

// 	conn.Write([]byte("+OK\r\n"))
// }

// func GetHandler(args []Value) Value {
// 	if len(args) != 1 {
// 		conn.Write([]byte("-ERR wrong number of arguments for 'get' command\r\n"))
// 		return
// 	}

// 	key := args[0]

// 	val, exist := StringStore.Get(key)
// 	if !exist {
// 		conn.Write([]byte("$-1\r\n"))
// 		return
// 	}

// 	reply := fmt.Sprintf("$%v\r\n%v\r\n", len(string(val)), val)
// 	conn.Write([]byte(reply))
// }
