package main

import (
	"first_app/chitchat/data"
	"net/http"
	"text/template"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0.:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	// files := []string{"templates/layout.html",
	// 	"templates/navbar.html",
	// 	"templates/index.html"}

	// templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		public_tmpl_files := []string{"templates/layout.html",
			"templates/navbar.html",
			"templates/index.html"}

		private_tmpl_files := []string{"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html"}

		var templates *template.Template

		if err != nil {
			// ParseFilesを使うとテンプレートファイルを解析してテンプレートを作れる
			templates = template.Must(template.ParseFiles(public_tmpl_files...))
		} else {
			templates = template.Must(template.ParseFiles(private_tmpl_files...))
		}
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
