package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
	"unsafe"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:7000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	start := time.Now()
	errorCount := 0

	for i := 0; i < 150000; i++ {
		data := createMessage(MessageTypeText, "Hello from client")
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("write error:", err)
			errorCount++
			if errorCount > 10 {
				fmt.Println("Too many errors, exiting loop")
				break
			}
		}
	}

	duration := time.Since(start)
	fmt.Println("duration time:", duration)

	// Optional: Sleep for a short period to ensure all data is sent before closing the connection
	time.Sleep(1 * time.Second)
}

const (
	MessageTypeJSON = 1
	MessageTypeText = 2
	MessageTypeXML  = 3
)

/*
0 1 2 3 | 4 5 6 7 | 8 N+
uint32  | uint32  | string
type    | length  | data
*/

func createMessage(mtype int, data string) []byte {
	buf := make([]byte, 4+4+len(data))
	binary.LittleEndian.PutUint32(buf[0:], uint32(mtype))
	binary.LittleEndian.PutUint32(buf[4:], uint32(len(data)))
	copy(buf[8:], []byte(data))
	return buf
}

func readMessage(data []byte) (mtype, mlen uint32, msg string) {
	mtype = binary.LittleEndian.Uint32(data[0:])
	mlen = binary.LittleEndian.Uint32(data[4:])
	//msg = string(data[8:])

	msgBytes := data[8:]
    msg = *(*string)(unsafe.Pointer(&msgBytes))
	return mtype, mlen, msg
}
