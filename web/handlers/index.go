package handlers

import (
	"net/http"

	"github.com/flosch/pongo2"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	template := pongo2.Must(pongo2.FromFile("web/templates/index.html"))
	context := pongo2.Context{}
	err := template.ExecuteWriter(context, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
}
