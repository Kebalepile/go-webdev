package main

import (
	"fmt"
	"html/template"
	_ "log"
	"net/http"
	"sync"
)

var (
	port string = ":8080"
	wg   sync.WaitGroup
	tpl  *template.Template
)

// type server bool

// func (x server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintln(w, "hello user")
// 	// fmt.Fprintln(w, r.Method)

// 	tpl.ExecuteTemplate(w, "index.html", "Keba")
// 	// wg.Done()
// }

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.html"))
}
func main() {
	// var handler server
	wg.Add(1)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	go http.ListenAndServe(port, nil)

	fmt.Println("server listening on server http://127.0.0.1", port)
	wg.Wait()
}


func homeHandler(w http.ResponseWriter, res *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", "Keba")
}
func aboutHandler(w http.ResponseWriter, res *http.Request) {

	tpl.ExecuteTemplate(w, "about.html", struct {
		Title, Msg string
	}{
		Title: "ABOUT",
		Msg:   "This is the about page.",
	})
}
