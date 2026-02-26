package main

import (
	"net/http"

	"gothink-course-project/01-intro/demoapp/pkg/stringutils"
)

func main() {
	http.ListenAndServe(
		":8080",
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				s := r.URL.Query()["s"]
				if len(s) == 0 {
					http.Error(w, "bad query", http.StatusBadRequest)
					return
				}

				rev := stringutils.Rev(s[0])

				w.Write([]byte(rev))
			},
		),
	)
}
