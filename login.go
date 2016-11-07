package main

import (
  "net/http"
  "fmt"
)

const (
    loginTop =
    `<!DOCTYPE HTML>
      <html>
        <head>
          <style>
            .error { color: red; }
          </style>
        </head>
        <title>Hello, It's Me</title>
        <body>`

    loginForm =
   `<form action="/session" method="POST">
      <label for="username">Username:</label>
      <input type="text" name="username"><br/ >
      <label for="password">Password:</label>
      <input type="password" name="password"><br/ >
      <input type="submit" value="Log In">
    </form>`
)

func login(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, loginTop, loginForm, pageBottom)
}
