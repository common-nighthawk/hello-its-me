package templates

func Style(style string) string {
  switch style {
  case "centered":
    return "body { text-align: center; } a { margin: 14px; }"
  case "error":
    return ".error { color: red; }"
  }
  panic ("invalid style option")
}
