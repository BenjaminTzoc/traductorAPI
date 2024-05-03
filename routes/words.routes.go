package routes

import (
	"encoding/json"
	"net/http"

	"github.com/BenjaminTzoc/TraductorAPI/db"
	"github.com/BenjaminTzoc/TraductorAPI/models"
	"github.com/gorilla/mux"
)

func GetWordsHandler(w http.ResponseWriter, r *http.Request) {
	var words []models.Word
	db.DB.Find(&words)
	json.NewEncoder(w).Encode(&words)
}

func GetWordHandler(w http.ResponseWriter, r *http.Request) {
	var word models.Word
	params := mux.Vars(r)

	db.DB.First(&word, params["id"])

	if word.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	json.NewEncoder(w).Encode(&word)
}

func PostWordsHandler(w http.ResponseWriter, r *http.Request) {
	var word models.Word
	json.NewDecoder(r.Body).Decode(&word)

	createdWord := db.DB.Create(&word)
	err := createdWord.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&word)
}

func DeleteWordsHandler(w http.ResponseWriter, r *http.Request) {
	var word models.Word
	params := mux.Vars(r)

	db.DB.First(&word, params["id"])

	if word.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Delete(&word)
	w.WriteHeader(http.StatusOK)
}
