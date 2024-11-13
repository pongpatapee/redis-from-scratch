// Seralizing and Deserializing RESP input/output

package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	STRING  = '+'
	ERROR   = '-'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'
)

type Value struct {
	typ   string
	str   string
	bulk  string
	array []Value
	num   int
}

type Resp struct {
	reader *bufio.Reader
}

func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

func (r *Resp) readLine() (line []byte, n int, err error) {
	// read until \r\n are the last two chracter
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}

		n += 1
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}

	}
	return line[:len(line)-2], n, nil
}

func (r *Resp) readInteger() (x int, n int, err error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}

	x, err = strconv.Atoi(string(line))
	if err != nil {
		return 0, n, err
	}

	return x, n, nil
}

func (r *Resp) Read() (Value, error) {
	_type, err := r.reader.ReadByte()
	if err != nil {
		return Value{}, err
	}

	switch _type {
	case ARRAY:
		return r.readArray()
	case BULK:
		return r.readBulk()
	default:
		fmt.Printf("Unknown type %v", string(_type))
		return Value{}, nil
	}
}

func (r *Resp) readArray() (Value, error) {
	v := Value{}
	v.typ = "array"

	length, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	v.array = make([]Value, length)
	for i := range length {
		val, err := r.Read()
		if err != nil {
			return v, err
		}

		v.array[i] = val
	}

	return v, nil
}

func (r *Resp) readBulk() (Value, error) {
	v := Value{}
	v.typ = "bulk"

	commandLength, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	bulk := make([]byte, commandLength)
	r.reader.Read(bulk)
	v.bulk = string(bulk)

	// parse and ignore CRLF
	r.readLine()

	return v, nil
}
