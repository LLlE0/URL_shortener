package handler

import (
	"github.com/LLlE0/URL_shortener/logger"
	store "github.com/LLlE0/URL_shortener/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"path"
)

func InitRoutes(s *store.Store) *chi.Mux {
	//init new router
	r := chi.NewRouter()
	// redirect /auth/ to /auth
	r.Use(middleware.RedirectSlashes)
	//seek for js in the 'js' folder
	fs := http.FileServer(http.Dir("../frontend/js/"))
	//seek for files all around the /frontend/ folder
	r.Handle("/*", fs)
	r.Get("/js/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Clean("../frontend"+r.URL.Path))
	})

	//serve all the api-routes

	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Get("/add", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		url := r.URL.Query().Get("url")
		if key == "" || url == "" {
			http.Error(w, "Missing key or url", http.StatusBadRequest)
			return
		}
		s.Save(key, url)
		log.Print(w, "Saved %s as %s", url, key)
	})

	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path[1:]
		url, ok := s.Load(key)
		if !ok {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}
		logger.Remember(r)
		http.Redirect(w, r, url, http.StatusFound)
	})

	return r
}
