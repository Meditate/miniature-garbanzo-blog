package main

import (
  "fmt"
  "net/http"
  "log"
  "html/template"

  "github.com/meditate/miniature-garbanzo-blog/models"
)

var posts map[string]*models.Post

func writeHandler(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles(
    "templates/write.html",
    "templates/header.html",
    "templates/footer.html")
  if err != nil {
    fmt.Println(w, err.Error())
  }
  t.ExecuteTemplate(w, "write", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles(
    "templates/index.html",
    "templates/header.html",
    "templates/footer.html")
  if err != nil {
    fmt.Println(w, err.Error())
  }

  fmt.Println(posts)

  t.ExecuteTemplate(w, "index", posts)
}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
  id := r.FormValue("id")
  title := r.FormValue("title")
  article := r.FormValue("article")

  post := models.NewPost(id, title, article)
  posts[post.Id] = post

  http.Redirect(w, r, "/", 302)
}

func main() {
  fmt.Println("Listen on port :3000")

  posts = make(map[string]*models.Post, 0)
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/write", writeHandler)
  http.HandleFunc("/SavePost", savePostHandler)

  err := http.ListenAndServe(":3000", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
