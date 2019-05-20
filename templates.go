package cms


import(
	"html/template"
	"time"
)

var Tmpl = template.Must(template.ParseGlob("../templates/*"))

type Page struct {
	Title string
	Content string
	Posts []*Post
}

type Comment struct {
	Author string
	Comment string
	DatePublished time.Time
}

type Post struct {
	Title string
	Content string
	DatePublished time.Time
	Comments []*Comment
}