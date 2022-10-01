package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//Moved code into type router and used as a mux
//func pathHandler(w http.ResponseWriter, r *http.Request) {
//switch r.URL.Path {
//case "/":
//homeHandler(w, r)
//case "/contact":
//contactHandler(w, r)
//default:
//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

//}

//}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>To get in touch email me at <ahref=\"mailto:test@gmail.com\">test@gmail.com<\a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "this is the faq page")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :30000")
	http.ListenAndServe(":3000", r)
}
