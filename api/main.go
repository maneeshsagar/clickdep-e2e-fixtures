package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	_ "github.com/lib/pq"
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
	var pg *sql.DB
	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		var err error
		pg, err = sql.Open("postgres", dsn)
		if err == nil {
			_, err = pg.Exec(`CREATE TABLE IF NOT EXISTS items (id SERIAL PRIMARY KEY, text TEXT)`)
		}
		if err != nil {
			log.Printf("db init: %v", err)
		}
	}
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) { js(w, map[string]any{"ok": true, "lang": "go"}) })
	http.HandleFunc("/api/env-echo", func(w http.ResponseWriter, r *http.Request) {
		host := ""
		if u, err := url.Parse(os.Getenv("DATABASE_URL")); err == nil {
			host = u.Host
		}
		js(w, map[string]any{"lang": "go", "env": e2e(), "has_database_url": os.Getenv("DATABASE_URL") != "", "database_host": host})
	})
	http.HandleFunc("/api/whoami", func(w http.ResponseWriter, r *http.Request) {
		names := []string{}
		for _, c := range r.Cookies() {
			names = append(names, c.Name)
		}
		has := false
		for _, n := range names {
			if n == "_gate_session" {
				has = true
			}
		}
		js(w, map[string]any{"lang": "go", "gate_email": r.Header.Get("X-Gate-Email"), "gate_app": r.Header.Get("X-Gate-App"), "cookie_names": names, "has_gate_session_cookie": has})
	})
	http.HandleFunc("/api/items", func(w http.ResponseWriter, r *http.Request) {
		if pg == nil {
			http.Error(w, `{"error":"no database"}`, 503)
			return
		}
		if r.Method == http.MethodPost {
			var b struct{ Text string `json:"text"` }
			json.NewDecoder(r.Body).Decode(&b)
			if _, err := pg.Exec(`INSERT INTO items(text) VALUES($1)`, b.Text); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
		rows, err := pg.Query(`SELECT text FROM items ORDER BY id`)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
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
