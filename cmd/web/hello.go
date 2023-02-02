// / greeting    greeting
// Welcome to my page.
package main

import (
	"log"
	"net/http"

	"github.com/MejiaFrancis/hello/handlers"
)

func main() {
	//create multiplexer
	mux := http.NewServeMux()
	//create file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) //exclude resource and go to static

	mux.HandleFunc("/greeting", handlers.Greeting) //passing in pointer, say where to find handler func
	// callback - above shows passing of the address not the func itself
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/message/create", handlers.MessageCreate)

	log.Println("Start server on port :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err) //should not reach here
}
