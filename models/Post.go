package models

type Post struct {
  Id      string
  Title   string
  Article string
}

func NewPost(id, title, article string) *Post {
  return &Post{ id, title, article }
}
