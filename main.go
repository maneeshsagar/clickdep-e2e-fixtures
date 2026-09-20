package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	_ "modernc.org/sqlite"
)

func e2e() map[string]string {
	m := map[string]string{}
	for _, kv := range os.Environ() {
		if strings.HasPrefix(kv, "E2E_") {
			p := strings.SplitN(kv, "=", 2)
			m[p[0]] = p[1]
		}
	}
	return m
}
func js(w http.ResponseWriter, v any) { w.Header().Set("Content-Type", "application/json"); json.NewEncoder(w).Encode(v) }

func main() {
	os.MkdirAll("/data", 0o755)
	db, err := sql.Open("sqlite", "/data/app.db")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS items (id INTEGER PRIMARY KEY AUTOINCREMENT, text TEXT)`); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1 id=ui>E2E-GO-SINGLE</h1><p>path: " + r.URL.Path + "</p>"))
	})
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) { js(w, map[string]any{"ok": true, "lang": "go-sqlite"}) })
	http.HandleFunc("/api/env-echo", func(w http.ResponseWriter, r *http.Request) {
		js(w, map[string]any{"lang": "go-sqlite", "env": e2e(), "has_database_url": os.Getenv("DATABASE_URL") != ""})
	})
	http.HandleFunc("/api/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var b struct{ Text string `json:"text"` }
			json.NewDecoder(r.Body).Decode(&b)
			if _, err := db.Exec(`INSERT INTO items(text) VALUES(?)`, b.Text); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
		rows, _ := db.Query(`SELECT text FROM items ORDER BY id`)
		out := []string{}
		for rows.Next() {
			var t string
			rows.Scan(&t)
			out = append(out, t)
		}
		js(w, map[string]any{"items": out})
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
