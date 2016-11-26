package templates

func Style(style string) string {
  switch style {
  case "centered":
    return `body { text-align: center; }
            a { margin: 14px; }
            a.figlet { margin: 0; text-decoration: none; color: inherit; }
            .error { color: red; }`
  case "error":
    return ".error { color: red; }"
  }
  panic ("invalid style option")
}
