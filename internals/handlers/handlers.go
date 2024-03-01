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
	"github.com/a-h/templ"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type response map[string]interface{}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	templ.Handler(templates.Home()).ServeHTTP()
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
		}
		http.SetCookie(w, &cookie)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response{
			"message": "success",
		})
	}
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
