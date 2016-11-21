package templates

import(
  "fmt"
  "../models"
)

const(
  Title = "Hello, It's Me"
  SignupForm =
   `<form action="/signup/create" method="POST">
      <label for="username">Username:</label>
      <input type="text" name="username"><br/ >
      <label for="password">Password:</label>
      <input type="password" name="password"><br/ >
      <label for="confirmation">Confirm Password:</label>
      <input type="password" name="confirmation"><br/ >
      <input type="submit" value="Sign Up">
    </form>`
  LoginForm =
    `<form action="/login/create" method="POST">
      <label for="username">Username:</label>
      <input type="text" name="username"><br/ >
      <label for="password">Password:</label>
      <input type="password" name="password"><br/ >
      <input type="submit" value="Log In">
    </form>`
)

func HTMLTop(style string) string {
  return fmt.Sprintf(`<!DOCTYPE HTML><html>
                      <head><style>%s</style></head>
                      <title>%s</title><body id='body'>`, style, Title)
}

func HTMLBottom() string {
  return `</body></html>`
}

func HTMLScript(script string) string {
  return fmt.Sprintf(`<script>%s</script>`, script)
}

func HTMLError(msg string) string {
  return fmt.Sprintf(`<p class="error">%s</p>`, msg)
}

func AudioPlayer(user *models.User, msg *models.Message) string {
  return fmt.Sprintf(`<audio controls>
                        <source src="assets/messages/%s/%s" type="audio/mpeg">
                      </audio>`, user.UUID, msg.File)
}
