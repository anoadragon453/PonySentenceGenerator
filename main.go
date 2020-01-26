package main

import (
	"fmt"
	"net/http"
	"log"
	"math/rand"
	"html/template"
	"time"
	"strconv"

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

func returnRandomSentence(w http.ResponseWriter, req *http.Request) {
	numPones := 3

	// Check if custom amount of ponies was defined
	if num := req.URL.Query().Get("ponies"); num != "" {
		newNum, err := strconv.Atoi(num)
		if err != nil {
			fmt.Fprint(w, "Error: 'ponies' query paramater not parsable int")
			return
		}
		numPones = newNum

		// No DDOS'ing!
		if numPones > 50 {
			fmt.Fprint(w, "Error: max 50 ponies")
			return
		}
		if numPones < 1 {
			fmt.Fprint(w, "Error: amount of ponies must be one or more")
			return
		}
	}
			
	// Print random sentence
	fmt.Fprint(w, ps.NewSentence(numPones))
}

func main() {
	http.HandleFunc("/ponysentence/sentence/", returnRandomSentence)
	http.HandleFunc("/ponysentence/sentence", returnRandomSentence)
	http.HandleFunc("/", generatePage)
	log.Print("Running server on :6969")
	err := http.ListenAndServe(":6969", nil)
	if err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
