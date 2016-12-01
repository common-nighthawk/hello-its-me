package templates

import(
  "../models"
  "fmt"
  "io"
  "github.com/common-nighthawk/go-figure"
)

const(
  Title = "Hello, It's Me"
  figletFont = "puffy"
  SignupForm =
   `<form action="/signup_create" method="POST">
      <label for="username">Username:</label>
      <input type="text" name="username"><br/ >
      <label for="password">Password:</label>
      <input type="password" name="password"><br/ >
      <label for="confirmation">Confirm Password:</label>
      <input type="password" name="confirmation"><br/ >
      <input type="submit" value="Sign Up">
    </form>`
  LoginForm =
    `<form action="/login_create" method="POST">
      <label for="username">Username:</label>
      <input type="text" name="username"><br/ >
      <label for="password">Password:</label>
      <input type="password" name="password"><br/ >
      <input type="submit" value="Log In">
    </form>`
  FindUserForm =
    `<form action="/message_new" method="GET">
       <label for="receiver_username">Username:</label>
       <input type="text" name="receiver_username">
       <input type="submit" value="Find User">
     </form>`
)

func HTMLTop(style string) string {
  return fmt.Sprintf(`<!DOCTYPE HTML><html>
                      <head>
                      <style>%s</style>
                      <script>%s</script>
                      </head>
                      <title>%s</title><body id='body'>`, style, gaScript, Title)
}

func HTMLBottom() string {
  return `</body></html>`
}

func WriteBanner(writer io.Writer, bannerTest string) {
  fmt.Fprint(writer, "<a class=\"figlet\" href=\"/actions\">")
  fmt.Fprint(writer, "<pre class=\"figlet\">")
  figure.Write(writer, figure.NewFigure(bannerTest, figletFont))
  fmt.Fprint(writer, "</pre>")
  fmt.Fprint(writer, "</a>")
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
