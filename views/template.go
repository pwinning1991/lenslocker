package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{
		htmlTpl: tpl,
	}, nil

}

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("error rendering the template: %v", err)
		http.Error(w, "there was an error rendering the template", http.StatusInternalServerError)
		return
	}

}
