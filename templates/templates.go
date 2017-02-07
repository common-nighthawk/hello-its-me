package templates

import(
  "fmt"
  "io"
  "github.com/common-nighthawk/go-figure"
)

type Args struct {
  StyleSheet string
  Script string
  ErrorMsg string
  File string
  UUID string
  Email string
  Timezone string
  ReceiverUsername string
}

const(
  figletFont = "puffy"
  ConfirmDelete = "Are you sure you want to delete this message? The action cannot be undone."
)

func WriteFigletBanner(writer io.Writer, bannerText string) {
  fmt.Fprint(writer, "<a class=\"figlet\" href=\"/actions\">")
  fmt.Fprint(writer, "<pre class=\"figlet\">")
  figure.Write(writer, figure.NewFigure(bannerText, figletFont, false))
  fmt.Fprint(writer, "</pre>")
  fmt.Fprint(writer, "</a>")
}

func WriteTextBanner(writer io.Writer, bannerText string) {
  fmt.Fprint(writer, "<div class='banner'><a href='/actions'>")
  fmt.Fprint(writer, bannerText)
  fmt.Fprint(writer, "</a></div>")
}
