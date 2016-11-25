package main

import(
  "./models"
  "fmt"
  "io"
  "net/http"
  "os"
  "time"
)

func message(w http.ResponseWriter, r *http.Request) {
  user, found := models.FindCurrentUser(r.Cookies(), db)
  //TODO: handle not found user gracefully
  if found == false { panic("user not found") }

  r.ParseMultipartForm(500)

  receiverUsername := r.FormValue("receiver_username")
  receiverUser, found := models.FindUserFromUsername(db, receiverUsername)
  if found == false { panic("user not found") }
  fmt.Println(receiverUsername)

  file := r.MultipartForm.File["blob"][0]
  outfileDir := fileDir(receiverUser)
  outfileName := fileName()

  infile, err := file.Open()
  err = os.Mkdir(outfileDir, os.ModePerm)
  outfile, err := os.Create(fmt.Sprintf("%s/%s", outfileDir, outfileName))
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }

  _, err = io.Copy(outfile, infile)
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }
  _, err = db.Exec("INSERT INTO messages VALUES($1, $2, $3)", user.Username, receiverUser.UUID, outfileName)
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }
  w.Write([]byte("success"))
}

func fileDir(user *models.User) string {
  return fmt.Sprintf("%s/messages/%s", fileServerDir, user.UUID)
}

func fileName() string {
  return time.Now().Format(time.RFC3339) + ".mp3"
}
