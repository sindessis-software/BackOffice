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

func Dashboard(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "dashboard.html", nil)
}

func Mostrador(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "mostrador.html", nil)
}

func Frutas(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "frutas.html", nil)
}

func Verduras(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "verduras.html", nil)
}

func Usuarios(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "usuarios.html", nil)
}

func Productos(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "productos.html", nil)
}

func Ventas(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "ventas.html", nil)
}

func Carrito(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "carrito.html", nil)
}

func Vendedor(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "vendedor.html", nil)
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
