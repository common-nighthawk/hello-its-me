package secrets

const(
  DBPassword = "1HA!YOUWISH"
  SSLCert = "/etc/letsencrypt/live/helloitsme.site/fullchain.pem"
  SSLKey = "/etc/letsencrypt/live/helloitsme.site/privkey.pem"
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
