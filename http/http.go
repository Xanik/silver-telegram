package http

import (
	"net/http"

	"go-challenge/http/handlers"
	"go-challenge/trigram"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/rs/cors"
)

func MountServer() *chi.Mux {
	router := chi.NewRouter()

	// Middlewares
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	// Create handlers
	handlersOpts := &handlers.HandlerOptions{
		Trigram: &trigram.TrigramIndex{},
	}
	httpHandlers := handlers.NewHTTPHandler(handlersOpts)

	// Add routes to handler
	router.Post("/learn", httpHandlers.HandleTeachRequest)
	router.Get("/generate", httpHandlers.HandleFetchRequest)

	// Health check
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.Data(w, r, []byte("Ok"))
	})

	return router
}
