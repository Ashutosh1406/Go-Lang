package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("processing the request")
	time.Sleep(8 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nhello world!\r\n"))
	conn.Close()
}
func main() {
	Listener, err := net.Listen("tcp", ":1729") //listen start
	if err != nil {
		log.Fatal(err)

	}
	var timing int = 0
	for {
		fmt.Println(timing)
		log.Println("waiting for a client to connect")
		conn, err := Listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("client connected")
		do(conn)
		timing++
	}

}
