package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pwinning1991/lenslocker/controllers"
	"github.com/pwinning1991/lenslocker/templates"
	"github.com/pwinning1991/lenslocker/views"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))
	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))
	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000")
	http.ListenAndServe("0.0.0.0:3000", r)
}
