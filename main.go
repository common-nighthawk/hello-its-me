package main

import (
  _ "github.com/lib/pq"
  "database/sql"
  "net/http"
  "log"
)

const (
    pageBottom =
        `</body>
      </html>`
)

var db *sql.DB

type User struct {
  username string
  password string
}

func init() {
  var err error
  db, err = sql.Open("postgres", "user=Daniel dbname=hello-its-me sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }

  if err = db.Ping(); err != nil {
    log.Fatal(err)
  }
}

func main() {
  http.HandleFunc("/", landing)
  http.HandleFunc("/signup", signup)
  http.HandleFunc("/create", create)
  http.HandleFunc("/login", login)
  http.ListenAndServe(":8080", nil)
}

