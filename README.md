# Redis from scratch

This is part of CodeCrafter's challenges on rebuilding popular tools from scratch.

## My progress

[![progress-banner](https://backend.codecrafters.io/progress/redis/d6baf7da-a325-4220-8aa9-8b2a53dc1722)](https://app.codecrafters.io/users/pongpatapee)

## Other resources

Built initial parsers and handlers on my own, but needed a more robust and structured approach.

[building redis from scratch](https://www.build-redis-from-scratch.dev/en/introduction)

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
