package main

import (
  "net/http"
  "fmt"
)

const (
    signupTop =
    `<!DOCTYPE HTML>
      <html>
        <head>
          <style>
          </style>
        </head>
        <title>Hello, It's Me</title>
        <body>`

    signupForm =
   `<form action="/create" method="POST">
      <label for="username">Username:</label>
      <input type="text" name="username"><br/ >
      <label for="password">Password:</label>
      <input type="password" name="password"><br/ >
      <label for="confirm">Confirm Password:</label>
      <input type="password" name="confirm"><br/ >
      <input type="submit" value="Sign Up">
    </form>`
)

func signup(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, signupTop, signupForm, pageBottom)
}
