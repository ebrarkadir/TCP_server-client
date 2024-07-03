package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime/pprof"
	"unsafe"
)

func main() {
	// CPU profil dosyası oluşturma
	f, err := os.Create("cpu.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// TCP bağlantı dinleme
	ls, err := net.Listen("tcp4", ":7000")
	if err != nil {
		panic(err)
	}
	defer ls.Close()

	fmt.Println("connection ready!")

	// Yeni bağlantıları kabul eden bir goroutine
	go func() {
		for {
			conn, err := ls.Accept()
			if err != nil {
				fmt.Println("Connection failed:", err)
				continue
			}
			go handler(conn) // Her bağlantıyı ayrı bir goroutine'de işleme
		}
	}()

	// Ctrl+C sinyalini yakalama ve programı düzgün kapatma
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func handler(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connection accepted:", conn.RemoteAddr().String())

	for {
		header := make([]byte, 8)
		_, err := conn.Read(header)
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		mlen := binary.LittleEndian.Uint32(header[4:])
		databuff := make([]byte, mlen)
		_, err = conn.Read(databuff)
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		var messagebuf []byte
		messagebuf = append(messagebuf, header...)
		messagebuf = append(messagebuf, databuff...)
		mtype, mlen, msg := readMessage(messagebuf)
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
	//msg = string(data[8:])

	msgBytes := data[8:]
    msg = *(*string)(unsafe.Pointer(&msgBytes))
	return mtype, mlen, msg
}
