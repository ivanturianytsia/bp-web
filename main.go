package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr = flag.String("addr", ":8081", "website address")
	flag.Parse()
	mux := http.NewServeMux()
	mux.Handle("/", &templateHandler{filename: "index.html"})
	mux.Handle("/new.html", &templateHandler{filename: "new.html"})
	mux.Handle("/view.html", &templateHandler{filename: "view.html"})
	log.Println("Seving website at: ", *addr)
	http.ListenAndServe(*addr, mux)
}
