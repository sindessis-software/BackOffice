package handlers

import (
	"html/template"
	"log"
	"net/http"
)

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("templates/login.html")
	template.Execute(rw, nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "index.html", nil)
}

func renderTemplate(w http.ResponseWriter, base, page string, data any) {
	tpl := template.Must(template.ParseFiles(base, templateDir+page))
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al renderizar platillas", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
