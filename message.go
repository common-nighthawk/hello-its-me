package main

import(
  "./models"
  "./secrets"
  "fmt"
  "io"
  "net/http"
  "os"
  "os/exec"
  "path/filepath"
  "strconv"
  "time"
)

func message(w http.ResponseWriter, r *http.Request) {
  currentUser, found := models.FindCurrentUser(r.Cookies(), db)
  if !found {
    http.Error(w, "no user logged in", 500)
    return
  }
  receiverUsername, explode := r.FormValue("receiver_username"), r.FormValue("explode")
  receiverUser, found := models.FindUserFromUsername(db, receiverUsername)
  explodeInSeconds, err := strconv.Atoi(explode)

  if !found || err != nil {
    http.Error(w, "invalid params: no user with receiver username and/or invalid explode value", 500)
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
  outfileName, err = convertWebMtoMP3(outfile)
  if err != nil {
    http.Error(w, "failed converting file on server", 500)
    return
  }

  explosionCol, explosionVal := explosionDetails(explodeInSeconds)
  insertStatement := fmt.Sprintf("INSERT into messages (sender_username, receiver_uuid, file, %s, created_at) VALUES ($1, $2, $3, $4, $5)", explosionCol)
  dbStatement, _ := db.Prepare(insertStatement)
  _, err = dbStatement.Exec(currentUser.Username, receiverUser.UUID, outfileName, explosionVal, time.Now().UTC())

  if err != nil {
    http.Error(w, "failed adding message to database", 500)
    return
  }

  if err = receiverUser.Notify(); err != nil {
    fmt.Println("failed sending notification email")
  }
}

func fileDir(user *models.User) string {
  return fmt.Sprintf("%s/%s", messagesDir, user.UUID)
}

func fileName() string {
  return time.Now().UTC().Format(time.RFC3339) + ".webm"
}

func convertWebMtoMP3(file *os.File) (string, error) {
  oldName := file.Name()
  newName := oldName[:len(oldName)-4] + "mp3"
  args := fmt.Sprintf("-i 'file:%s' -vn -ab 128k -ar 44100 -y 'file:%s'",  oldName, newName)
  err := exec.Command("bash", "-c", secrets.FFmpeg(env()) + args).Run()
  return filepath.Base(newName), err
}

func explosionDetails(explodeInSeconds int) (attr string, value interface{}) {
  attr, value = "explode_after", -explodeInSeconds
  if explodeInSeconds > 0 {
    attr, value = "expires_at", time.Now().UTC().Add(time.Duration(explodeInSeconds) * time.Second)
  }
  return attr, value
}
