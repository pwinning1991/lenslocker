package main

import (
	"fmt"
	"github.com/gorilla/csrf"
	"github.com/pwinning1991/lenslocker/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/pwinning1991/lenslocker/controllers"
	"github.com/pwinning1991/lenslocker/templates"
	"github.com/pwinning1991/lenslocker/views"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//TODO maybe look at moving this to the controllers somehow
	r.Get("/metrics", promhttp.Handler().ServeHTTP)
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))
	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))
	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err) //TODO handle this error better
	}
	defer db.Close()
	userService := models.UserService{db}
	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))
	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"signin.gohtml", "tailwind.gohtml",
	))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/users/me", usersC.CurrentUser)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000")
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMW := csrf.Protect(
		[]byte(csrfKey),
		//TODO Fix this before deploying
		csrf.Secure(false))
	http.ListenAndServe("0.0.0.0:3000", csrfMW(r))
}
