package main

import "testing"

func TestMessageParse(t *testing.T) {
	message := "hello"
	data := createMessage(MessageTypeText, message)
	mtype, mlen, msg := readMessage(data)
	if mtype != MessageTypeText {
		t.Errorf("Invalid message type. got: %d, want: %d", mtype, MessageTypeText)
	}

	if mlen != uint32(len(message)) {
		t.Errorf("Invalid length. got: %d, want: %d", mlen, len(message))
	}

	if msg != message {
		t.Errorf("Invalid message. got: %s, want: %s", msg, message)
	}
}
