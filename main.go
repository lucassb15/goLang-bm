package main

import (
	"fmt"
	"html/template"
	"log"
	"modulo/src/model/form"
	"modulo/src/model/index"
	"net/http"
	"os"
)

// type usuario struct {
// 	Titulo    string
// 	Nome      string
// 	Sobrenome string
// 	Email     string
// 	Idade     int
// 	Success   bool
// }

var tpl *template.Template // ponteiro do pacote

func main() {
	port := os.Getenv("PORT_WEB")                      // env port
	fmt.Println("Server started on localhost: ", port) // msg
	tpl = template.Must(tpl.ParseGlob("src/view/templates/*.html"))
	http.HandleFunc("/", index.Index)
	http.HandleFunc("/form", form.Form)
	log.Fatal(http.ListenAndServe(port, nil))
}

// // INDEX
// func index(w http.ResponseWriter, r *http.Request) { // Solicitação ao servidor
// 	tpl.ExecuteTemplate(w, "index.html", nil)
// }

// // FORM
// func form(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		tpl.Execute(w, nil)
// 		return
// 	}
// 	userform := usuario{
// 		Nome:  r.FormValue("nome"),
// 		Email: r.FormValue("email"),
// 	}
// 	tpl.Execute(w, struct {
// 		Success  bool
// 		Userform usuario
// 	}{true, userform})

// }
