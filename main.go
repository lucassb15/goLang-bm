package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type usuario struct {
	Titulo    string
	Nome      string
	Sobrenome string
	Email     string
	Idade     int
	Success   bool
}

func (me *usuario) calculaIdade() { // calcular idade usando info do usuario

}

var tpl *template.Template
var user1 usuario

func usuarioText() {
	user1 = usuario{
		Titulo:    "Usuário 1",
		Nome:      "Lucas",
		Sobrenome: "Barbosa",
		Idade:     22,
		Email:     "lucas@exemplo.com",
		Success:   false,
	}
}

func main() {
	usuarioText()
	port := os.Getenv("PORT_WEB") // env port
	fmt.Println("Server started on localhost: ", port)
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/", index)
	http.HandleFunc("/form", form)
	log.Fatal(http.ListenAndServe(port, nil))

}

// INDEX
func index(w http.ResponseWriter, r *http.Request) { // Solicitação ao servidor

	tpl.ExecuteTemplate(w, "index.html", user1)
}

// FORM
func form(w http.ResponseWriter, r *http.Request) {

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
