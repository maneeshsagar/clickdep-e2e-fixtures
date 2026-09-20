package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if os.Getenv("E2E_REQUIRED") == "" { log.Fatal("FATAL: E2E_REQUIRED is not set") }
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "TINY-OK f9-requires-env") })
	log.Printf("listening on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
