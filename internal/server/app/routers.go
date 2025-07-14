package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"main/internal/server/app/middlewares"
	"time"
)

// NewRouters constructs and configures the main router with middleware and routes.
func NewRouters(h *Handlers, conf middlewares.AuthParams) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middlewares.Authentication(conf))

	r.Use(middleware.Timeout(60 * time.Second))

	//r.Route("/", func(r chi.Router) {
	//	r.Get("/ping", h.health.Ping)
	//	r.Post("/", h.links.AddLinkInText)
	//	r.Route("/{id}", func(r chi.Router) {
	//		r.Get("/", h.links.GetLink)
	//	})
	//	r.Route("/api", func(r chi.Router) {
	//		r.Route("/user", func(r chi.Router) {
	//			r.Get("/urls", h.users.GetLinks)
	//			r.Delete("/urls", h.users.DeleteLinks)
	//		})
	//		r.Route("/shorten", func(r chi.Router) {
	//			r.Post("/", h.links.AddLink)
	//			r.Post("/batch", h.links.AddLinks)
	//		})
	//		r.Route("/internal", func(r chi.Router) {
	//			r.Get("/stats", h.stats.GetMainStats)
	//		})
	//	})
	//})
	return r
}
