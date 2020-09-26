package handlers

import (
	"fmt"
	"net/http"

	"github.com/forum/models"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// PostThread POST /thread/post.
func PostThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "thread_not_found",
			})
			errorMessage(writer, request, msg)
		}
		user, err := sess.User()
		if err != nil {
			fmt.Println("Cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := models.ThreadByUUID(uuid)
		if err != nil {
			fmt.Println("Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			fmt.Println("Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
