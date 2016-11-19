package main

import (
  "fmt"
  "database/sql"
  "log"
  "net/http"
  _ "github.com/lib/pq"
  "os"
  // "mime/multipart"
  "io"
  "strconv"
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

func savefile(w http.ResponseWriter, r *http.Request) {
  r.ParseMultipartForm(500) 
  fmt.Println(r.MultipartForm.File["key"])

  for _, fheaders := range r.MultipartForm.File {
    for _, hdr := range fheaders {
      infile, _ := hdr.Open()
      fmt.Println("infile")
      fmt.Println(infile)
      outfile, err := os.Create("./assets/messages/" + hdr.Filename)
      fmt.Println("outfile")
      fmt.Println(err)
      fmt.Println(outfile)
      written, _ := io.Copy(outfile, infile)
      fmt.Println("written")
      fmt.Println(written)
      w.Write([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))))
      fmt.Println("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written)))
      fmt.Println("x")
    }
  }

  http.Error(w, http.StatusText(500), 500)
  return
}
