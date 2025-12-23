package handlers

import (
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

func Login(rw http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("templates/login.html")
	template.Execute(rw, nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	p, err := getUsuario(1)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	renderTemplate(w, templateBase, "index.html", p)
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

func getUsuario(id int) (*User, error) {
	url := fmt.Sprintf("http://localhost:3000//users/%d", id)
	fmt.Println("Obteniendo usuario con ID:", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error al hacer la petición GET: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("la API devolvió un código de estado: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	var user User
	if err := decoder.Decode(&user); err != nil {
		fmt.Printf("error al decodificar JSON: %w", err)
	}
	return &user, nil
}

type User struct {
	User      string `json:"User"`
	UserName  string `json:"UserName"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Acceso    []Acceso
}

type Acceso struct {
	TypeAccess uint   `json:"type_access"`
	DescAccess string `json:"description_access"`
	RollAccess uint   `json:"roll_access"`
	Status     bool   `json:"status"`
	UserID     uint   `json:"user_id"`
}
