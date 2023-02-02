package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// TERMINAL
func helloworld(name string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Escreva seu nome para abrir a web: ")
	scanner.Scan()
	nome := scanner.Text()
	fmt.Printf("Hello, %q", nome)
}

// WEB
func index(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprint(w, "Página não encontrada!")
	} else {
		template.Execute(w, "ADMIN")
	}
}

func main() {
	helloworld("")
	fmt.Printf("\n Server started on localhost:8000 ")
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
