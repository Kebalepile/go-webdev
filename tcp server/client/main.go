package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}
	// close once everyting is done or error occurs
	defer conn.Close()

	n, err := conn.Write([]byte("I called you... little boy"))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("number of bytes wrote:", n)

}
