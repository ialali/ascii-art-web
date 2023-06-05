package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type PageData struct {
	asciiartcontainer string
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/generate", handleGenerate)
	// http.HandleFunc("/fonts/", handleFonts)

	fmt.Println("Server started. Listening on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func renderTemplate(w http.ResponseWriter, templateFile string, data interface{}) {
	tmpl := template.Must(template.ParseFiles(templateFile))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	font := r.FormValue("font")
	result := generateAsciiArt(text, font)
	data := PageData{
		asciiartcontainer: result,
	}
	renderTemplate(w, "index.html", data)
}

func handleFonts(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func generateAsciiArt(text, font string) string {
	words := strings.Split(text, "\n")
	rawBytes, err := os.ReadFile(fmt.Sprintf("fonts/%s.txt", font))
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.ReplaceAll(string(rawBytes), "\r\n", "\n"), "\n")
	var result strings.Builder
	for i, word := range words {
		if word == "" {
			if i < len(words)-1 {
				result.WriteString(word)
			}
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

		return result.String()
	}
	return result.String()
}
