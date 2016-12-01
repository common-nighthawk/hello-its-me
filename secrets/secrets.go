package secrets

const(
  DBPassword = "HA!YOUWISH"
  GACode = "HA!YOUWISH"
  SSLCert = "/home/ubuntu/.ssl/cert.pem"
  SSLKey = "/home/ubuntu/.ssl/key.pem"
)

func FileServerDir(env string) string {
  if env == "prod" {
    return "/home/ubuntu/go/src/hello-its-me/assets"
  }
  return "/Users/Daniel/Documents/go-workspace/src/hello-its-me/assets"
}
