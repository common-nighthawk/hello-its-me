package main

import (
  "net/http"
)

const (
    pageBottom =
        `</body>
      </html>`
)

func main() {
  http.HandleFunc("/", landing)
  http.HandleFunc("/signup", signup)
  http.HandleFunc("/login", login)
  http.ListenAndServe(":8080", nil)
}
