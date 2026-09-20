package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	_ = os.Getenv
	http.HandleFunc("/" THIS IS NOT GO, func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "TINY-OK f11-bad-dockerfile") })
	log.Printf("listening on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
