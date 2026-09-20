package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	_ = os.Getenv
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "TINY-OK f12-loopback-bind") })
	log.Printf("listening on 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
