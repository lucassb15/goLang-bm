package index

import (
	"html/template"
	"net/http"
)

var tpl *template.Template // ponteiro do pacote

// INDEX
func Index(w http.ResponseWriter, r *http.Request) { // Solicitação ao servidor
	tpl.ExecuteTemplate(w, "index.html", nil)
}
