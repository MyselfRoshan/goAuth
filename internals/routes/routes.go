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

	// Serve the static files from the ./public folder to localhost:<port-number>/
	// Note: always use ./<folder-name> or <folder-name> instead of /<folder-name> to serve files
	fileServer := http.FileServer(http.Dir("public"))
	router.Handle("/*", http.StripPrefix("/", fileServer))

	// Handle Routing
	router.Get("/", handlers.HandleIndex)

	router.Route("/register", func(router chi.Router) {
		router.Get("/", handlers.HandleRegisterPage)
		router.Post("/", handlers.HandleRegister)
	})

	router.Route("/login", func(router chi.Router) {
		router.Get("/", handlers.HandleLoginPage)
		router.Post("/", handlers.HandleLogin)
	})

	router.Get("/dashboard", handlers.HandleDashboard)
	router.Post("/logout", handlers.HandleLogout)
	return router
}
