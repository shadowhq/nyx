package nyx

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	GetRepo("git://github.com/gopheracademy/gopheracademy-web.git")

	fmt.Fprintf(w, "nyx API")
}

func Commits(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "commits")
}
