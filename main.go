package main

import (
	"fmt"
	"net/http"
	"log"
	"math/rand"
	"html/template"
	"time"

	ps "github.com/anoadragon453/ponysentence"
)

type Page struct {
	Sentence string
}

func generatePage(w http.ResponseWriter, req *http.Request) {
	// Generate a sentence with 1-3 ponies
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sentence := ps.NewSentence(r.Intn(3) + 1)

	// Make the webpage
	page := Page{sentence}

	t, err := template.ParseFiles("page.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to generate page: %s", err.Error())
		return
	}
	t.Execute(w, page)
}

func main() {
	http.HandleFunc("/", generatePage)
	log.Print("Running server on :6969")
	err := http.ListenAndServe(":6969", nil)
	if err != nil {
		log.Fatal("Unable to start server: ", err)
	}

}
