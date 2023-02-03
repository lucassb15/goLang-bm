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
	port := os.Getenv("PORT_WEB")
	fmt.Println("Server started on localhost: ", port)
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/", index)
	http.HandleFunc("/form", form)
	log.Fatal(http.ListenAndServe(port, nil))
}

// WEB
func index(w http.ResponseWriter, r *http.Request) { // Solicitação ao servidor

	tpl.ExecuteTemplate(w, "index.html", user1)
}

func form(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/form.html")
	case "POST":
		fmt.Fprint(w, "Post: \n", r.PostForm)
		Nome := r.FormValue("nome")
		Email := r.FormValue("email")

		fmt.Fprintf(w, "Nome: %s\n", Nome)
		fmt.Fprintf(w, "Email: %s\n", Email)

	default:
		fmt.Fprintf(w, "Pega apenas o post")
	}

}
