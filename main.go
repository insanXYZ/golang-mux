package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/insanXYZ/golang-mux/config"
	"github.com/insanXYZ/golang-mux/controller"
	"net/http"
)

func main() {
	mux := config.Mux()
	sql := config.InitDB()
	contactsController := controller.NewContactsController(sql)

	mux.HandleFunc("GET /contacts", contactsController.GetAllContacts)
	mux.HandleFunc("POST /contacts", contactsController.Insert)
	mux.HandleFunc("DELETE /contacts/", contactsController.Delete)
	mux.HandleFunc("PUT /contacts/", contactsController.Update)

	server := http.Server{
		Handler: mux,
		Addr:    "localhost:1323",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
