package main

import (
	"fmt"
	"net"
)

func main() {

	ls, err := net.Listen("tcp4", ":7000")
	if err != nil {
		panic(err)	
	}

	fmt.Println("connection ready!")

	for{
		conn, err := ls.Accept()
		if err != nil {
			fmt.Println("Connection failed:", err)
			continue
		}
		handler(conn)
	}
}

func handler(conn net.Conn){
	fmt.Println("Connection accepted:", conn.RemoteAddr().String())

	for{
		buf := make([]byte, 8)
		_, err := conn.Read(buf[:])//read bufferından data okuyor.
		if err != nil {
			fmt.Println("Read error:", err)
			conn.Close()
			break
		}

		fmt.Printf("message client: %s\n", buf)

		_, err = conn.Write(buf)
		if err != nil {
			fmt.Println("Write error:", err)
			continue
		}
	}
}