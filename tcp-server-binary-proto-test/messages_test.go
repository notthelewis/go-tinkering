package main

import (
	"testing"
)

// This could potentially lose some valuable stack trace information...
func run_invalid_length_check(buf []byte, t *testing.T) {
	msg, e := parse_hello(buf)

	if e == nil {
		t.Error("No error thrown when length requirement was unmet")
	}

	if msg.protocol_version != 0 || msg.software_version != 0 || msg.content_length != 0 || msg.message_type != 0 {
		t.Error("Struct still constructed when length check should not have passed")
	}
}

func Test_parse_hello_length_check(t *testing.T) {
	emptyBuffer := []byte{}
	run_invalid_length_check(emptyBuffer, t)

	oneByteUnder := []byte{
		0x00,
		0x01,
		0x10,
		0x11,
		0x01,
		// MISSING CONTENT_LENGTH
	}

	run_invalid_length_check(oneByteUnder, t)
}

func Test_parse_hello_valid(t *testing.T) {
	helloMessage := []byte{
		0x00,0x01,
		0x00,0x01,
		0x02,
		0x01,
	}

	msg, e := parse_hello(helloMessage)
	if e != nil {
		t.Error(e.Error())
	}

	expectedMessage := HelloMessage{
		protocol_version: 1,
		software_version: 1,
		message_type:     COMMAND,
		content_length:   1,
	}

	if msg != expectedMessage {
		t.Errorf("Expected message incorrect, expected: %v, received: %v", expectedMessage, msg)
	}

}
