package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/shurcooL/home/httputil"
	"github.com/shurcooL/home/internal/page/resume"
	"github.com/shurcooL/httperror"
	"github.com/shurcooL/notifications"
	"github.com/shurcooL/reactions"
	"github.com/shurcooL/users"
)

var resumeHTML = template.Must(template.New("").Funcs(template.FuncMap{"noescape": func(s string) template.HTML { return template.HTML(s) }}).Parse(`<html>
	<head>
		<title>Dmitri Shuralyov - Resume</title>
		<link href="/icon.png" rel="icon" type="image/png">
		<meta name="viewport" content="width=device-width">
		<link href="/assets/fonts/fonts.css" rel="stylesheet" type="text/css">
		<link href="/assets/resume/style.css" rel="stylesheet" type="text/css">

		{{noescape "<!-- Unminified source is at https://github.com/shurcooL/resume. -->"}}
		<script async src="/assets/resume/resume.js"></script>

		{{if .Production}}` + googleAnalytics + `{{end}}
	</head>
	<body>`))

const googleAnalytics = `<script>
		  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
		  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
		  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
		  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

		  ga('create', 'UA-56541369-3', 'auto');
		  ga('send', 'pageview');

		</script>`

func initResume(reactions reactions.Service, notifications notifications.Service, usersService users.Service) {
	http.Handle("/resume", cookieAuth{httputil.ErrorHandler(usersService, func(w http.ResponseWriter, req *http.Request) error {
		if req.Method != "GET" {
			return httperror.Method{Allowed: []string{"GET"}}
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		data := struct{ Production bool }{*productionFlag}
		err := resumeHTML.Execute(w, data)
		if err != nil {
			return err
		}

		// Optional (still experimental) server-side rendering.
		prerender, _ := strconv.ParseBool(req.URL.Query().Get("prerender"))
		if prerender {
			authenticatedUser, err := usersService.GetAuthenticated(req.Context())
			if err != nil {
				log.Println(err)
				authenticatedUser = users.User{} // THINK: Should it be a fatal error or not? What about on frontend vs backend?
			}
			returnURL := req.RequestURI
			err = resume.RenderBodyInnerHTML(req.Context(), w, reactions, notifications, usersService, time.Now(), authenticatedUser, returnURL)
			if err != nil {
				return err
			}
		}

		_, err = io.WriteString(w, `</body></html>`)
		return err
	})})
}
