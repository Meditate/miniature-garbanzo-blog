package main

import (
  "fmt"
  "net/http"
  "log"
  "html/template"

  "github.com/meditate/miniature-garbanzo-blog/models"

  "database/sql"
  _ "github.com/lib/pq"
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

func editHandler(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles(
    "templates/write.html",
    "templates/header.html",
    "templates/footer.html")
  if err != nil {
    fmt.Println(w, err.Error())
  }

  id := r.FormValue("id")

  fmt.Println(id)
  post, found := posts[id]

  if !found {
    http.NotFound(w, r)
  }

  t.ExecuteTemplate(w, "write", post)
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

	var post *models.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.Article = article
	} else {
		id = GenerateId()
		post := models.NewPost(id, title, article)
		posts[post.Id] = post
	}

	http.Redirect(w, r, "/", 302)
}

func destroyPostHandler(w http.ResponseWriter, r *http.Request) {
  id := r.FormValue("id")

  if id != "" {
    delete(posts, id)
  }

  http.Redirect(w, r, "/", 302)
}

func newSessionHandler(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles(
    "templates/sessions/new.html",
    "templates/header.html",
    "templates/footer.html")
  if err != nil {
    fmt.Println(w, err.Error())
  }

  t.ExecuteTemplate(w, "sessions/new", nil)
}

func main() {
  fmt.Println("Listen on port :3000")

  db, db_err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
  if db_err != nil {
    log.Fatal(db_err)
  }

  rows, _ := db.Query("SELECT * FROM articles")
  if rows != nil {
    log.Fatal(rows)
  }

  posts = make(map[string]*models.Post, 0)
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/write", writeHandler)
  http.HandleFunc("/edit", editHandler)
  http.HandleFunc("/destroy", destroyPostHandler)
  http.HandleFunc("/SavePost", savePostHandler)
  http.HandleFunc("/sessions/new", newSessionHandler)

  err := http.ListenAndServe(":3000", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
