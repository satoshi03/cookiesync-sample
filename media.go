package main

import (
	"net/http"
	"text/template"
)

type Page struct {
	CookieValue string
	CookieKey   string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	cookie := http.Cookie{
		Name:  "mediaCookieID",
		Value: "mediaCookieID",
		Path:  "/",
	}
	http.SetCookie(w, &cookie)

	page := Page{cookie.Name, cookie.Value}
	tmpl, err := template.ParseFiles("index.html") // ParseFilesを使う
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler) // hello
	http.ListenAndServe(":8080", nil)
}
