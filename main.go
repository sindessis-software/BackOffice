package main

import (
	"log"
	"net/http"

	"github.com/sindessis-software/BackOffice/handlers"
)

func main() {
	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.Index)
	/*router.HandleFunc("/login", handlers.Login)
	router.HandleFunc("/usuarios", handlers.Usuarios)
	router.HandleFunc("/dashboard", handlers.Dashboard)
	router.HandleFunc("/inventarios", handlers.Inventarios)
	router.HandleFunc("/about", handlers.About)
	router.HandleFunc("/productos", handlers.Productos)
	router.HandleFunc("/contacto", handlers.Contacto)*/
	port := ":8090"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
