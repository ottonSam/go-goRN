package main

import (
  "net/http"
  "go-goRN/internal"
)

func main() {
  handler := internal.NewHandler()
  http.ListenAndServe(":8080", handler)
}