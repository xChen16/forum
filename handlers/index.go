package handlers

import (
	"net/http"

	"github.com/forum/models"
)

// Index 论坛首页路由处理器方法.
func Index(writer http.ResponseWriter, request *http.Request) {
	//files := []string{"views/layout.html", "views/navbar.html", "views/index.html"}
	//templates := template.Must(template.ParseFiles(files...))
	threads, err := models.Threads()
	if err == nil {
		//templates.ExecuteTemplate(w, "layout", threads)
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "auth.navbar", "index")
		}
	}
}
