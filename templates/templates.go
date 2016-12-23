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
  ReceiverUsername string
}

const figletFont = "puffy"

func WriteBanner(writer io.Writer, bannerTest string) {
  fmt.Fprint(writer, "<a class=\"figlet\" href=\"/actions\">")
  fmt.Fprint(writer, "<pre class=\"figlet\">")
  figure.Write(writer, figure.NewFigure(bannerTest, figletFont, false))
  fmt.Fprint(writer, "</pre>")
  fmt.Fprint(writer, "</a>")
}
