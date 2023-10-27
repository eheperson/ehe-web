package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var appEnv string

func init() {
	file, err := os.Open("config.conf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 && parts[0] == "APP_ENV" {
			appEnv = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles("templates/base.html", tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, p)
}

type Page struct {
	Env string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/home.html", &Page{Env: appEnv})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/about.html", &Page{Env: appEnv})
}

func artHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/art.html", &Page{Env: appEnv})
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/art", artHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
