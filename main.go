package main

import (
  "fmt"
  "database/sql"
  "log"
  "net/http"
  "os"
  "runtime"
  _ "github.com/lib/pq"
)

const dbname = "hello-its-me"
var db *sql.DB
var fileServerDir string = os.Getenv("GOPATH") + "src/hello-its-me/assets"

func init() {
  var err error
  db, err = sql.Open("postgres", dbSource(env()))
  if err != nil { log.Fatal(err) }
  if err = db.Ping(); err != nil {
    log.Fatal(err)
  }
}

func main() {
  http.HandleFunc("/", landing)
  http.HandleFunc("/signup", signup)
  http.HandleFunc("/signup_create", createSignup)
  http.HandleFunc("/login", login)
  http.HandleFunc("/login_create", createLogin)
  http.HandleFunc("/actions", actions)
  http.HandleFunc("/message_new", newMessage)
  http.HandleFunc("/message", message)
  http.HandleFunc("/messages", messages)
  http.HandleFunc("/assets/", assets)

  if env() == "prod" {
    go http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      secureURL := fmt.Sprintf("https://%s%s", r.Host, r.URL.String())
      http.Redirect(w, r, secureURL, http.StatusMovedPermanently)
    }))
    http.ListenAndServeTLS(":443", "/home/ubuntu/.ssl/cert.pem", "/home/ubuntu/.ssl/key.pem", nil)
  } else {
    http.ListenAndServe(":8080", nil)
  }
}

func env() string {
  if runtime.GOOS == "linux" {
    return "prod"
  } else if runtime.GOOS == "darwin" {
    return "dev"
  }
  panic("program not running on mac or linux")
}

func dbSource(env string) string {
  if env == "prod" {
    dbuser, dbpassword := "postgres", "HA!YOUWISH"
    return fmt.Sprintf("postgres://%s:%s@localhost/%s", dbuser, dbpassword, dbname)
  }
  dbuser, sslmode := "Daniel", "disable"
  return fmt.Sprintf("user=%s dbname=%s sslmode=%s", dbuser, dbname, sslmode)
}
