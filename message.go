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
  currentUser, found := models.FindCurrentUser(r.Cookies(), db)
  if !found {
    http.Error(w, "no user logged in", 500)
    return
  }
  receiverUsername := r.FormValue("receiver_username")
  receiverUser, found := models.FindUserFromUsername(db, receiverUsername)
  if !found {
    http.Error(w, "no user with receiver username", 500)
    return
  }

  //TODO determine and update maxMemory with more appropriate value
  r.ParseMultipartForm(500)
  file := r.MultipartForm.File["blob"][0]
  outfileDir := fileDir(receiverUser)
  outfileName := fileName()
  infile, err := file.Open()
  err = os.Mkdir(outfileDir, os.ModePerm)
  outfile, err := os.Create(fmt.Sprintf("%s/%s", outfileDir, outfileName))
  _, err = io.Copy(outfile, infile)
  if err != nil {
    http.Error(w, "failed saving file to server", 500)
    return
  }

  _, err = db.Exec("INSERT INTO messages VALUES($1, $2, $3)", currentUser.Username, receiverUser.UUID, outfileName)
  if err != nil {
    http.Error(w, "failed adding message to database", 500)
    return
  }
}

func fileDir(user *models.User) string {
  return fmt.Sprintf("%s/messages/%s", fileServerDir, user.UUID)
}

func fileName() string {
  return time.Now().Format(time.RFC3339) + ".mp3"
}
