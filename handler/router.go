package handler

import (
	"github.com/LLlE0/URL_shortener/logger"
	store "github.com/LLlE0/URL_shortener/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"html/template"
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

		log.Print(r.URL)

		url := r.URL.Query().Get("url")

		t, _ := template.ParseFiles("../index.html")
		var key string
		var err error
		if url != "" {
			key, err = s.Save(url)
			if err != nil {
				log.Print(err)
			}
		}

		//Template to return the link
		t.Execute(w, struct {
			K string
		}{
			K: key,
		})

		log.Print("Saved " + url + " as " + key)
	})

	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path[1:]
		if key == "" {
			http.Redirect(w, r, "/add", http.StatusFound)
		}
		url, ok := s.Load(key)
		if !ok {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}
		logger.Remember(r)
		log.Print("Redirecting from " + key + " to " + url)
		http.Redirect(w, r, url, http.StatusFound)
	})

	return r
}
