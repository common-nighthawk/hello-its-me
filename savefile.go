package main

import(
  "./models"
  "fmt"
  "io"
  "net/http"
  "os"
  "time"
)

func savefile(w http.ResponseWriter, r *http.Request) {
  user, found := models.FindUser(r.Cookies(), db)
  //TODO: handle not found user gracefully
  if found == false { panic("user not found") }

  r.ParseMultipartForm(500)
  file := r.MultipartForm.File["blob"][0]
  infile, _ := file.Open()
  outfileDir := fmt.Sprintf("%s/messages/%s", fileServerDir, user.UUID)
  os.Mkdir(outfileDir, os.ModePerm)
  fileName := time.Now().Format(time.RFC3339) + ".mp3"
  outfile, _ := os.Create(outfileDir + "/" + fileName)
  written, _ := io.Copy(outfile, infile)
  fmt.Fprint(w, "success" + string(written))
  // result, err := db.Exec("INSERT INTO books VALUES($1, $2, $3)", user.Username, "aaab", "aaab")
  db.Exec("INSERT INTO messages VALUES($1, $2, $3)", user.Username, "aaab", "aaab/" + fileName)
  fmt.Fprint(w, "good")

  // for _, fheaders := range r.MultipartForm.File {
  //   for _, hdr := range fheaders {
  //     infile, _ := hdr.Open()
  //     outfile, err := os.Create("./assets/messages/" + hdr.Filename + time.Now().String())
  //     written, _ := io.Copy(outfile, infile)
  //     w.Write([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))))
  //   }
  // }

}
