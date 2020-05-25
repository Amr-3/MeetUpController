package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("port", ":8010", "http service address") // Q=17, R=18

func main() {
	flag.Parse()
	http.Handle("/", nil)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

