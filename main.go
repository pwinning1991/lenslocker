package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pwinning1991/lenslocker/controllers"
	"github.com/pwinning1991/lenslocker/views"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(tpl))
	tpl = views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(tpl))
	tpl = views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controllers.StaticHandler(tpl))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000")
	http.ListenAndServe("0.0.0.0:3000", r)
}
