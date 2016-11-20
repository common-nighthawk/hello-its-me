package main

import (
  "fmt"
  "database/sql"
  "log"
  "net/http"
  _ "github.com/lib/pq"
)

const(
  dbuser = "Daniel"
  dbname = "hello-its-me"
  sslmode = "disable"
  fileServerDir = "/Users/Daniel/Documents/go-workspace/src/hello-its-me/assets"
)

var db *sql.DB

func init() {
  var err error
  dbConf := fmt.Sprintf("user=%s dbname=%s sslmode=%s", dbuser, dbname, sslmode)
  db, err = sql.Open("postgres", dbConf)
  if err != nil { log.Fatal(err) }
  if err = db.Ping(); err != nil {
    log.Fatal(err)
  }
}

func main() {
  http.HandleFunc("/", landing)
  http.HandleFunc("/signup", signup)
  http.HandleFunc("/create", create)
  http.HandleFunc("/login", login)
  http.HandleFunc("/session", session)
  http.HandleFunc("/actions", actions)
  http.HandleFunc("/messages", messages)
  http.HandleFunc("/send", send)
  http.HandleFunc("/savefile", savefile)

  fileServer := http.FileServer(http.Dir(fileServerDir))
  http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

  http.ListenAndServe(":8080", nil)
}
