package main

import (
	"bufio"
	"fmt"
	_ "io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	// close TCP server once done with main go function
	defer li.Close()

	// infinite while loop, handles one requst at atime
	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}
		// writes back as a respnse
		// io.WriteString(conn, "\n Hello from Tcp Server \n")
		// fmt.Fprintln(conn, "How's your day?")
		// fmt.Fprintf(conn, "%v", "Well, I hope.")
		// writes back as a respnse
		// go address(conn.LocalAddr(), conn.RemoteAddr())
		// handle multiple connections to TCP server
		go handler(conn)
		// conn.Close()

	}

	// go routine for each request
	// go tcpServer()

}

// func address(local, remote net.Addr) {

// 	fmt.Println("local address is", local, "remote address is", remote)
// }

func handler(conn net.Conn) {

	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Println(conn, "I heard you say:\n", ln)
	}

}
