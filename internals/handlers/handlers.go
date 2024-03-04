package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MyselfRoshan/goAuth/internals/database"
	"github.com/MyselfRoshan/goAuth/internals/models"
	"github.com/MyselfRoshan/goAuth/templates/pages"
	"github.com/golang-jwt/jwt/v5"
)

type response map[string]interface{}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	pages.Index().Render(r.Context(), w)
}

func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	SECRET := []byte(os.Getenv("JWT_SRCRET"))
	if err != nil {
		log.Fatal("Error getting the cookie")
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Panic(err)
		json.NewEncoder(w).Encode(response{
			"message": "unauthenticated",
		})
	}
	// get claims from token.Claims and type cast by using .(*jwt.RegisteredClaims)
	claims := token.Claims.(*jwt.RegisteredClaims)
	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	json.NewEncoder(w).Encode(user)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		HttpOnly: true,
		Expires:  time.Now().Add(-time.Hour),
	}

	http.SetCookie(w, &cookie)
	json.NewEncoder(w).Encode(response{
		"message": "sucessfully logged out",
	})
}
