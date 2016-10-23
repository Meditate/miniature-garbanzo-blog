package main

import (
  "fmt"
  "net/http"
  "log"
  "html/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles(
    "templates/index.html",
    "templates/header.html",
    "templates/footer.html")
  if err != nil {
    fmt.Println(w, err.Error())
  }
  t.ExecuteTemplate(w, "index", nil)
}

func main() {
  fmt.Println("Listen on port :3000")

  http.HandleFunc("/", indexHandler)

  err := http.ListenAndServe(":3000", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
