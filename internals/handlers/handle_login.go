package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/MyselfRoshan/goAuth/internals/database"
	"github.com/MyselfRoshan/goAuth/internals/models"
	"github.com/MyselfRoshan/goAuth/templates/pages"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HandleLoginPage(w http.ResponseWriter, r *http.Request) {
	pages.Login().Render(r.Context(), w)
}
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	var user models.User
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&data)

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{
			"message": "incorrect password",
		})
	} else {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.Id)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		})

		// Always convert secret to byte before using it inside claims.SignedString()
		SECRET := []byte(os.Getenv("JWT_SRCRET"))
		token, err := claims.SignedString(SECRET)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
			json.NewEncoder(w).Encode(response{
				"messsage": "could not login",
			})
		}
		// save token to cookie
		cookie := http.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, &cookie)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response{
			"message": "success",
		})
	}
}
