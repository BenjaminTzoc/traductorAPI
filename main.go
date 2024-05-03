package main

import (
	"net/http"

	"github.com/BenjaminTzoc/TraductorAPI/db"
	"github.com/BenjaminTzoc/TraductorAPI/models"
	"github.com/BenjaminTzoc/TraductorAPI/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Word{})
	db.DB.AutoMigrate(models.Language{})
	db.DB.AutoMigrate(models.Category{})
	db.DB.AutoMigrate(models.WordCategory{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/words", routes.GetWordsHandler).Methods("GET")
	r.HandleFunc("/words/{id}", routes.GetWordHandler).Methods("GET")
	r.HandleFunc("/words", routes.PostWordsHandler).Methods("POST")
	r.HandleFunc("/words/{id}", routes.DeleteWordsHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
