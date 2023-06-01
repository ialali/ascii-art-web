package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/generate", handleGenerate)
	http.HandleFunc("/fonts/", handleFonts)

	fmt.Println("Server started. Listening on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Parse the JSON data
	var data struct {
		Word string `json:"word"`
		Font string `json:"font"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate the ASCII art
	result, err := generateAsciiArt(data.Word, data.Font)
	if err != nil {
		http.Error(w, "Failed to generate ASCII art", http.StatusInternalServerError)
		return
	}

	// Write the result as the response
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}

func handleFonts(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func generateAsciiArt(word string, font string) (string, error) {
	if word == "" {
		return "", fmt.Errorf("Please enter a word")
	}

	filename := "fonts/standard.txt"
	switch font {
	case "shadow.txt":
		filename = "fonts/shadow.txt"
	case "thinkertoy.txt":
		filename = "fonts/thinkertoy.txt"
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Failed to read font file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	var result strings.Builder

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

	return result.String(), nil
}
