package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

var user User

func Login(rw http.ResponseWriter, r *http.Request) {
	restartValue()
	template, _ := template.ParseFiles("templates/login.html")
	template.Execute(rw, nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
		}
		user.Usuario_email = r.Form.Get("email")
		user.Usuario_password = r.Form.Get("password")
		p, err := validaUsuario(user)
		if err == nil && p != nil {
			renderTemplate(w, templateBase, "index.html", p)
		} else {
			fmt.Println("Error :", err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
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

func Paquetes(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "paquetes.html", nil)
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

func validaUsuario(user User) (*User, error) {
	url := "http://localhost:3000/usuariosValida"
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("error al serializar el usuario: %w", err)
		return nil, err
	}
	fmt.Println("Struct as JSON:", string(jsonData))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("error al hacer la petición POST: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("la API devolvió un código de estado: %d", resp.StatusCode)
		return nil, nil
	}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&user); err != nil {
		fmt.Printf("error al decodificar JSON: %w", err)
	}
	return &user, nil
}

type User struct {
	Usuario_nombre    string `gorm:"not null" json:"nombre"`
	Usuario_apellido1 string `gorm:"not null" json:"apellido1"`
	Usuario_apellido2 string `gorm:"not null" json:"apellido2"`
	Usuario_password  string `gorm:"not null" json:"password"`
	Usuario_email     string `gorm:"not null;unique_index" json:"email"`
	Usuario_permisos  string `gorm:"not null" json:"permisos"`
	Usuario_activo    string `gorm:"not null" json:"activo"`
}

func restartValue() {
	user = User{}
}
