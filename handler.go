package cms

import(
	"net/http"
	"strings"
	"time"
)

// Servepage serves a page based on the route matched. This will match any URL
//Beginning with /page. WOW

func servePage (w http.ResponseWriter, r *http.Request){
	path := strings.TrimLeft(r.URL.Path, "/page/")

	if path == "" {
		http.NotFound(w,r)
		return
	}

	p := &Page{
		Title: strings.ToTitle(path),
		Content:"Here is my page",
	}
	Tmpl.ExecuteTemplate(w, "Page", p)
}

func ServePost(w http.ResponseWriter, r *http.Request){
	path  := strings.TrimLeft(r.URL.Path, "/post/")

	if path == "" {
		http.NotFound(w,r)
		return
	}

	p := &Post{
		Title: strings.ToTitle(path),
		Content: "HEre is my page",
		Comments : []*Comment{
			&Comment{
				Author: "Malcom Gladwell",
				Comment: "Outlier",
				DatePublished: time.Now(),
			},
		},
	}
	Tmpl.ExecuteTemplate(w, "post", p)
}




func HandleNew(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		Tmpl.ExecuteTemplate(w, "new", nil)

	case "POST":
		title := r.FormValue("tittle")
		content := r.FormValue("content")
		contentType := r.FormValue("content-type")
		r.ParseForm()

		if contentType == "page"{
			Tmpl.ExecuteTemplate(w,"page", &Page{
				Title: title,
				Content: content,
			})
			return
		}

		if contentType == "post" {
			Tmpl.ExecuteTemplate(w, "post", &Post{
				Title: title,
				Content: content,
			})
			return
		}
	default:
		http.Error(w, "Method not supported: " + r.Method, http.StatusMethodNotAllowed)

	}
}
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "GO projects CMS",
		Content: "Welcome to our home page!",
		Posts:[] *Post{
		&Post{
		Title: "Hello, world",
		Content: "Thanks for visiting this page",
		DatePublished: time.Now(),
		},
		&Post{
		Title: "A post with Comments ",
		Content: "Here's a controversial post. It's a sure to attract comments",
		DatePublished: time.Now().Add(-time.Hour),
		Comments: [] *Comment{
			&Comment{
			Author: "ELon Musk",
			Comment: "NeverMind, I was kidding",
			DatePublished: time.Now().Add(-time.Hour / 2),
		},
	},
},
},
}
	Tmpl.ExecuteTemplate(w, "Page", p)
}
