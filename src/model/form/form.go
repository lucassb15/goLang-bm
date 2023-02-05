package form

import (
	"html/template"
	"net/http"
)

type usuario struct {
	Titulo    string
	Nome      string
	Sobrenome string
	Email     string
	Idade     int
	Success   bool
}

var tpl *template.Template // ponteiro do pacote

// FORM
func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tpl.Execute(w, nil)
		return
	}
	userform := usuario{
		Nome:  r.FormValue("nome"),
		Email: r.FormValue("email"),
	}
	tpl.Execute(w, struct {
		Success  bool
		Userform usuario
	}{true, userform})

}
