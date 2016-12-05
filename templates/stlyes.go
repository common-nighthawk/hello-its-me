package templates

func Style(style string) string {
  switch style {
  case "centered":
    return `body { text-align: center; }
            a { margin: 14px; }
            a.figlet { margin: 0; text-decoration: none; color: inherit; }
            button { margin: 0 5px; }
            .error { color: red; }
            .error a { margin: 0; }
            #start, #stop, #dismiss, #send { display: none; }
            #rec { display: none; color: red; animation: blinker 1.5s linear infinite; }
            @keyframes blinker {  50% { opacity: 0; }}`
  case "error":
    return ".error { color: red; }"
  }
  panic ("invalid style option")
}
