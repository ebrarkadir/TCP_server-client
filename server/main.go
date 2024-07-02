package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	ls, err := net.Listen("tcp4", ":7000")
	if err != nil {
		panic(err)
	}

	fmt.Println("connection ready!")

	for {
		conn, err := ls.Accept()
		if err != nil {
			fmt.Println("Connection failed:", err)
			continue
		}
		go handler(conn) // Her bağlantıyı ayrı bir goroutine'de işlemek için go anahtar kelimesi kullanılır.
	}
}

func handler(conn net.Conn) {
	fmt.Println("Connection accepted:", conn.RemoteAddr().String())

	for {
		header := make([]byte, 8)
		_, err := conn.Read(header) // read bufferından data okuyor.
		if err != nil {
			fmt.Println("Read error:", err)
			conn.Close()
			break
		}

		mlen := binary.LittleEndian.Uint32(header[4:])
		databuff := make([]byte, mlen)
		_, err = conn.Read(databuff)
		if err != nil {
			fmt.Println("Read error:", err)
			conn.Close()
			break
		}

		var messagebuf []byte
		messagebuf = append(messagebuf, header...)
		messagebuf = append(messagebuf, databuff...)
		mtype, _, msg := readMessage(messagebuf)
		fmt.Printf("type: %d, len: %d, msg: %s\n", mtype, mlen, msg)
	}
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
	msg = string(data[8:])
	return mtype, mlen, msg
}
