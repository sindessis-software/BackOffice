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

	router.HandleFunc("/", handlers.Login)
	router.HandleFunc("/Index", handlers.Index)
	router.HandleFunc("/Ventas", handlers.Ventas)
	router.HandleFunc("/Productos", handlers.Productos)
	router.HandleFunc("/Usuarios", handlers.Usuarios)
	router.HandleFunc("/Carrito", handlers.Carrito)
	router.HandleFunc("/Mostrador", handlers.Mostrador)
	router.HandleFunc("/Dashboard", handlers.Dashboard)
	router.HandleFunc("/Frutas", handlers.Frutas)
	router.HandleFunc("/Paquetes", handlers.Paquetes)
	router.HandleFunc("/Verduras", handlers.Verduras)
	router.HandleFunc("/Vendedor", handlers.Vendedor)
	port := ":8090"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
