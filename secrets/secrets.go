package secrets

const(
  DBPassword = "HA!YOUWISH"
  SSLCert = "/home/ubuntu/.ssl/cert.pem"
  SSLKey = "/home/ubuntu/.ssl/key.pem"
)

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
