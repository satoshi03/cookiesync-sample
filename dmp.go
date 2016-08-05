package main

import (
	"fmt"
	"net/http"
)

func cookieSyncHandler(w http.ResponseWriter, r *http.Request) {

	mCID := r.FormValue("cookie_id")

	cookie := http.Cookie{
		Name:  "gunosyCookieName",
		Value: "gunosyCookieValue",
		Path:  "/",
	}
	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "image/gif")
	fmt.Println(mCID)
}

func main() {
	http.HandleFunc("/cs.gif", cookieSyncHandler) // hello
	http.ListenAndServe(":8081", nil)
}
