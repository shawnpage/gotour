package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting String
	Punct    String
	Who      String
}

func (p String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, p)
}

func (p Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, p.Greeting, p.Punct, p.Who)
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed not."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
