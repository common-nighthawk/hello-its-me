package templates

import(
  "../models"
  "fmt"
  "io"
  "github.com/common-nighthawk/go-figure"
)

const figletFont = "puffy"

type Args struct {
  StyleSheet string
}

func WriteBanner(writer io.Writer, bannerTest string) {
  fmt.Fprint(writer, "<a class=\"figlet\" href=\"/actions\">")
  fmt.Fprint(writer, "<pre class=\"figlet\">")
  figure.Write(writer, figure.NewFigure(bannerTest, figletFont, false))
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
  return fmt.Sprintf(`<audio controls id=%s>
                        <source src="assets/messages/%s/%s" type="audio/mpeg">
                      </audio>`, msg.File, user.UUID, msg.File)
}
