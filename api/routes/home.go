package routes

import (
	"net/http"
	"os"
)

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("templates/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(content)
}
