# Redis from scratch

[![progress-banner](https://backend.codecrafters.io/progress/redis/d6baf7da-a325-4220-8aa9-8b2a53dc1722)](https://app.codecrafters.io/users/pongpatapee)

This is part of CodeCrafter's challenges on rebuilding popular tools from scratch.

## Concepts used

- TCP
- Concurrency

  - Goroutines
  - Timers
  - Mutex locks

- Redis Serialization Protocal (RESP)

## How it works

### Overview

Redis clients communicates with a Redis server through the _Redis Serialization Protocol RESP_.
The Redis server communicates through the network layer and listens for TCP connections from the clients.

RESP is not TCP specific but defines the format of the incoming payload. Typically it starts the string
off with a character `*`, `+`, `-`, etc. to define the data type of the information being received.
Each value is delimited with `\r\n`

Here are some data type examples:

- `+` -> Simple strings
- `-` -> Simple Errors
- `$` -> Bulk strings
- `*` -> Arrays

A command from the client could look like this

```
*2\r\n
$5\r\n
hello\r\n
$5\r\n
world\r\n
```

Building the Redis server means that we are able to properly _parse_ commands in RESP,
_execute_ these commands, and _respond_ back to the client in RESP.

### Concurrency

#### Original

According to Redis (the company), Redis is mostly single threaded. This means that
Redis is able to serve all requests from multiple clients using a single thread.
They use a technique called **multiplexing**, which enables them to serve
requests sequentially while not being slow (Similar to how Node.js is designed).

They are still able to process requests extremely quickly despite, the single thread
because Redis is designed to not block system calls, e.g., I/O to a socket.

Although since Redis 2.4, they use _some_ threads in the background for slow I/O
operation to disk, but is still serving requests through a single thread.

#### For this project

Since I'm building the server in Go, it makes sense to use one of Go's main feature -- Goroutines.

Each server connection from the client, I am handling their requests through a goroutine.

```go
 for {
  conn, err := listener.Accept() // blocking call
  if err != nil {
   fmt.Println("Error accepting connection: ", err.Error())
   continue
  }

  go handleRequest(conn)
 }
```

### Command parsing and Marshalling

### Data storage

#### In-memory storage

#### Data persistence

## Other resources

Built initial parsers and handlers on my own, but needed a more robust and structured approach.
So I followed the command parsing and marshalling from
[this article](https://www.build-redis-from-scratch.dev/en/introduction).
