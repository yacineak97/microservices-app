package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		loadAndRenderTemplate(writer, "test.page.gohtml")
	})

	fmt.Println("Front-end service is starting on port 8000")
	if startErr := http.ListenAndServe(":8000", nil); startErr != nil {
		log.Fatalf("Server failed to start: %v", startErr)
	}
}

func loadAndRenderTemplate(writer http.ResponseWriter, fileName string) {
	baseTemplates := []string{
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/header.partial.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
	}

	templatesToLoad := append([]string{fmt.Sprintf("./cmd/web/templates/%s", fileName)}, baseTemplates...)

	parsedTemplates, parseErr := template.ParseFiles(templatesToLoad...)
	if parseErr != nil {
		log.Printf("Error parsing templates: %v", parseErr)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	executionErr := parsedTemplates.Execute(writer, nil)
	if executionErr != nil {
		log.Printf("Error executing templates: %v", executionErr)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}
