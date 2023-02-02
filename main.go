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
	scanner := bufio.NewScanner(os.Stdin) // ler input
	fmt.Printf("Escreva seu nome para abrir a web: ")
	scanner.Scan() // armazena input
	nome := scanner.Text()
	fmt.Printf("Hello, %q", nome) // output
}

// WEB
func index(w http.ResponseWriter, r *http.Request) { // Solicitação ao servidor
	template, err := template.ParseFiles("templates/index.html") // leitura do arquivo que foi passado

	if err != nil {
		fmt.Fprint(w, "Página não encontrada!")
	} else {
		template.Execute(w, "ADMIN") // executa e passa para o html o texto "admin"
	}
}

func main() {
	helloworld("")                                     // chamando a função
	fmt.Printf("\n Server started on localhost:8000 ") // apenas uma msg que o servidor foi iniciado
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
