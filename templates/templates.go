package templates

import(
  "fmt"
  "../models"
)

const(
  SignupForm =
   `<form action="/create" method="POST">
      <label for="username">Username:</label>
      <input type="text" name="username"><br/ >
      <label for="password">Password:</label>
      <input type="password" name="password"><br/ >
      <label for="confirmation">Confirm Password:</label>
      <input type="password" name="confirmation"><br/ >
      <input type="submit" value="Sign Up">
    </form>`
  LoginForm =
    `<form action="/session" method="POST">
      <label for="username">Username:</label>
      <input type="text" name="username"><br/ >
      <label for="password">Password:</label>
      <input type="password" name="password"><br/ >
      <input type="submit" value="Log In">
    </form>`
)

func HTMLTop(style string) string {
  return fmt.Sprintf(`<!DOCTYPE HTML><html><head><style>%s</style></head>
                      <title>%s</title><body>`, style, "Hello, It's Me")
}

func HTMLBottom() string {
  return `</body></html>`
}

func HTMLError(msg string) string {
  return fmt.Sprintf(`<p class="error">%s</p>`, msg)
}

func AudioPlayer(msg *models.Message) string {
  return fmt.Sprintf(`<audio controls>
                        <source src="assets/messages/%s" type="audio/mpeg">
                      </audio>`, msg.Path)
}

func Style(style string) string {
  switch style {
  case "centered":
    return "body { text-align: center; } a { margin: 14px; }"
  case "error":
    return ".error { color: red; }"
  }
  panic ("invalid style option")
}
