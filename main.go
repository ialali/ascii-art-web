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
	fmt.Println("Server is live on http://localhost:3030 . To turn off the server use comand 'control + C'")
	http.HandleFunc("/", handleMainPage)
	http.HandleFunc("/ascii-art", handleAsciiArt)
	http.ListenAndServe("localhost:3030", nil)
}

func handleMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // Set 200 status explicitly
	err = tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("font")
	if strings.Contains(text, "\r\n") {
		text = strings.ReplaceAll(text, "\r\n", "\\n")
	}

	result := generateAsciiArt(text, banner)
	pagedata := &PageData{AsciiArt: result}
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, pagedata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
