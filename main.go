package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type PageData struct {
	AsciiArt string
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handleMainPage)
	http.HandleFunc("/ascii-art", handleAsciiArt)
	fmt.Println("Server started. Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found ", http.StatusNotFound)
		return
	}
	renderTemplate(w, "index.html", nil)
}

func handleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("font")

	result := generateAsciiArt(text, banner)

	data := PageData{
		AsciiArt: result,
	}

	renderTemplate(w, "index.html", data)
}

func renderTemplate(w http.ResponseWriter, templateFile string, data interface{}) {
	tmpl := template.Must(template.ParseFiles(templateFile))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func generateAsciiArt(text, banner string) string {
	words := strings.Split(text, `\n`)
	rawBytes, err := os.ReadFile(fmt.Sprintf("fonts/%s.txt", banner))
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.ReplaceAll(string(rawBytes), "\r\n", "\n"), "\n")
	var result strings.Builder

	for _, word := range words {
		if word == "" {
			result.WriteString("\n")
			continue
		}

		for h := 1; h < 9; h++ {
			for _, l := range word {
				for lineIndex, line := range lines {
					if lineIndex == (int(l)-32)*9+h {
						result.WriteString(line)
					}
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}
