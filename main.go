package main

import (
  "./secrets"
  "database/sql"
  "fmt"
  "log"
  "net/http"
  "html/template"
  "runtime"
  _ "github.com/lib/pq"
)

const dbname = "hello-its-me"
var db *sql.DB
var publicDir string = secrets.PublicDir(env())
var messagesDir string = secrets.MessagesDir(env())

var templateHTMLTop = findTemplate("html-top")
var templateHTMLBottom = findTemplate("html-bottom")
var templateScript = findTemplate("script")
var templateErrorMsg = findTemplate("error-message")

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
  http.HandleFunc("/logout", logout)
  http.HandleFunc("/actions", actions)
  http.HandleFunc("/message_new", newMessage)
  http.HandleFunc("/message", message)
  http.HandleFunc("/message_update", updateMessage)
  http.HandleFunc("/messages", messages)
  http.HandleFunc("/assets/", assets)

  fileServer := http.FileServer(http.Dir(publicDir))
  http.Handle("/public/", http.StripPrefix("/public/", fileServer))

  if env() == "prod" {
    go http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      secureURL := fmt.Sprintf("https://%s%s", r.Host, r.URL.String())
      http.Redirect(w, r, secureURL, http.StatusMovedPermanently)
    }))
    http.ListenAndServeTLS(":443", secrets.SSLCert, secrets.SSLKey, nil)
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
    dbuser, dbpassword := "postgres", secrets.DBPassword
    return fmt.Sprintf("postgres://%s:%s@localhost/%s", dbuser, dbpassword, dbname)
  }
  dbuser, sslmode := "Daniel", "disable"
  return fmt.Sprintf("user=%s dbname=%s sslmode=%s", dbuser, dbname, sslmode)
}

func findTemplate(name string) *template.Template {
  file := fmt.Sprintf("%s/templates/%s.html", publicDir, name)
  template, err := template.ParseFiles(file)
  if err != nil {
    panic("invalid template")
  }
  return template
}
