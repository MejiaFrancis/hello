package handlers

import (
	"fmt"
	"net/http"
	"time"
)

// include --about
// include --home
// create handler for greeting
func Greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my page."))
}

// create handler for about
func About(w http.ResponseWriter, r *http.Request) {
	day := time.Now().Weekday()
	w.Write([]byte(fmt.Sprintf("Welcome to my  about page, have a nice %s", day)))
}

// create handler for home
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my home page."))
}

// create handler for home
func MessageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("ALLOW", "POST") //setting header in order to do a 'write'
		//w.WriteHeader(405) //write in header
		http.Error(w, "Method Not Allowed, http.StatusMethodNotAllowed", 500)
		//w.Write([]byte("method not allowed")) //this is writing in the body
		return
	}
	w.Write([]byte("method created..."))

}
