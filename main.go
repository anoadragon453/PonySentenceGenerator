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

type Image struct {
	URL string
}

type Page struct {
	Sentence string
	Images []Image
}

// generatePage generates a pony-themed sentence and returns a webpage with the
// sentence and associated images on it.
func generatePage(w http.ResponseWriter, req *http.Request) {
	// Generate a sentence with 1-3 ponies
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sentence, imageURLs := ps.NewSentenceWithImages(r.Intn(3) + 1)
	fmt.Println(sentence, imageURLs)

	// Grab all of the returned images
	var images []Image
	for _, url := range imageURLs {
		images = append(images, Image{
			URL: url,
		})
	}

	// Make a page containing our sentence and images
	page := Page{sentence, images}

	// Parse the custom template HTML file 
	t, err := template.ParseFiles("page.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to generate page: %s", err.Error())
		return
	}
	// Fill in the template with our sentence and images
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
