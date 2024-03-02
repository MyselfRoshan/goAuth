package routes

import (
	"net/http"

	"github.com/MyselfRoshan/goAuth/internals/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		// To use JWT token
		AllowCredentials: true,
		MaxAge:           300,
	}))

	fileServer := http.FileServer(http.Dir("/static"))
	router.Handle("/static/*", http.StripPrefix("static", fileServer))

	// router.Route("/", func(router chi.Router) {
	router.Get("/", handlers.HandleIndex)
	router.Post("/register", handlers.HandleRegister)
	router.Post("/login", handlers.HandleLogin)
	router.Get("/dashboard", handlers.HandleDashboard)
	router.Post("/logout", handlers.HandleLogout)
	// })
	return router
}
