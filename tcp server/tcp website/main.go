package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net"
	"os"
	"strings"
)

var (
	tpl   *template.Template
	err   interface{}
	pages = make(map[string]func(conn net.Conn, uri, method string))
)

func init() {
	tpl, err = template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatal("shit happended,", err)
	}

	pages["home"] = func(conn net.Conn, uri, method string) {
		if m := strings.ToUpper(method); m != "GET" {

			fmt.Fprint(conn, "HTTP/1.1 405 Method Not Allowed \r\n")
			// fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
			fmt.Fprint(conn, "Content-Type: text/html\r\n")
			fmt.Fprint(conn, "\r\n")
			getTemplate(conn, "405.gohtml", struct {
				Title, Method, URI string
			}{
				Title:  "405",
				Method: m,
				URI:    uri,
			})
		} else if m := strings.ToUpper(method); m == "GET" {
			getTemplate(conn, "index.gohtml", "Home Page")
		}

	}

	pages["about"] = func(conn net.Conn, uri, method string) {
		if m := strings.ToUpper(method); m != "GET" {

			fmt.Fprint(conn, "HTTP/1.1 405 Method Not Allowed \r\n")
			// fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
			fmt.Fprint(conn, "Content-Type: text/html\r\n")
			fmt.Fprint(conn, "\r\n")
			getTemplate(conn, "405.gohtml", struct {
				Title, Method, URI string
			}{
				Title:  "405",
				Method: m,
				URI:    uri,
			})
		} else if m := strings.ToUpper(method); m == "GET" {
			getTemplate(conn, "about.gohtml", "About Page")
		}
	}
}

func main() {

	// creates a tcp server
	listener, err := net.Listen("tcp", ":8080")
	fmt.Println("TCP server running on localhost:8080")
	defer listener.Close()

	if err != nil {
		log.Fatalln(err.Error())
	}

	for {
		conn, err := listener.Accept() //accept any incoming connection

		defer conn.Close()

		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handler(conn) // go routine for ever connection recived

	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	request(conn)

}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		} else if ln == "" {
			// headers are done
			break
		}
		i++
	}

	fmt.Println("Number of lines read in headers,", i)
}
func mux(conn net.Conn, ln string) {

	m := strings.Fields(ln)[0] // get  request method
	u := strings.Fields(ln)[1] // get uri
	fmt.Println(u)
	// router
	switch u {
	default:
		pages["home"](conn, u, m) // conn net.Conn, uri, method string

	case "/about":
		pages["about"](conn, u, m) // conn net.Conn, uri, method string

	}

	fmt.Print("\n------------------***** end of TCP server response *****------------------ \n")
}

func getTemplate(conn net.Conn, tplName string, data interface{}) {
	// does not print response on dial server cmd
	err = tpl.ExecuteTemplate(os.Stdout, tplName, data)
	if err != nil {
		log.Fatalln("occurde while executing template", err)
	}
}
