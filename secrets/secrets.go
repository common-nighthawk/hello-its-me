package secrets

const(
  DBPassword = "1HA!YOUWISH"
  SSLCert = "/home/ubuntu/.ssl/cert.pem"
  SSLKey = "/home/ubuntu/.ssl/key.pem"
  Email = "helloitsmewebsite@gmail.com"
  SMTPPassword = "2HA!YOUWISH"
)

func PublicDir(env string) string {
  if env == "prod" {
    return "/home/ubuntu/go/src/hello-its-me/public"
  }
  return "/Users/Daniel/Documents/go-workspace/src/hello-its-me/public"
}

func MessagesDir(env string) string {
  if env == "prod" {
    return "/home/ubuntu/messages"
  }
  return "/Users/Daniel/Documents/go-workspace/src/hello-its-me/messages"
}

func FFmpeg(env string) string {
  if env == "prod" {
    return "/usr/bin/ffmpeg "
  }
  return "/usr/local/bin/ffmpeg "
}
