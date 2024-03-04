package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MyselfRoshan/goAuth/internals/database"
	"github.com/MyselfRoshan/goAuth/internals/models"
	"github.com/MyselfRoshan/goAuth/templates/pages"
	"golang.org/x/crypto/bcrypt"
)

func HandleRegisterPage(w http.ResponseWriter, r *http.Request) {
	pages.Register().Render(r.Context(), w)
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&data)

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: string(password),
	}

	database.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}
