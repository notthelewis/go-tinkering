package main

import (
	"errors"
	"fmt"
)

type MessageType uint8

const (
	// HACK: Skip zero as this is what Go initializes struct fields as by default
	// kinda odd behaviour considering there's a 'nil' type

	ERROR   MessageType = 0x01
	COMMAND MessageType = 0x02
)

// TODO: learn module system and split each message into individual file in a 'Messages' directory
type HelloMessage struct {
	protocol_version uint16
	software_version uint16
	message_type     MessageType
	content_length   uint8
}

// TODO: Work out how to get size of struct without constant
const HELLO_MSG_LEN_IN_BYTES = 6

// type == 0x01
type CommandMessage struct {
	command uint8
	params  uint8
}

const COMMAND_MSG_LEN_IN_BYTES = 2

// TODO: Look into streaming parser...
func parse_hello(buffer []byte) (HelloMessage, error) {
	if len(buffer) < HELLO_MSG_LEN_IN_BYTES {
		retMsg := HelloMessage{}
		return retMsg, errors.New(fmt.Sprintf("error::parse_hello::Buffer length invalid. Expected = %q, received = %q",
			HELLO_MSG_LEN_IN_BYTES,
			len(buffer)))
	}

	messageType := get_message_type(buffer[4])
	// TODO: Handle more error cases

	return HelloMessage{
		protocol_version: uint16(buffer[0])<<8 | uint16(buffer[1]),
		software_version: uint16(buffer[2])<<8 | uint16(buffer[3]),
		message_type:     messageType,
		content_length:   buffer[5],
	}, nil
}

func get_message_type(t uint8) MessageType {
	switch t {
	case 2:
		return COMMAND
	default:
		return ERROR
	}
}
