package controllers

import (
	"html/template"
	"net/http"

	"github.com/pwinning1991/lenslocker/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes we offer a free trial for 30 days",
		},
		{
			Question: "What are your suppot hours?",
			Answer:   "We have support staff ansering emails 24/7",
		},
		{
			Question: "How do i contact support?",
			Answer:   `Email us - <a href="mailto:suppport@test.com">support@test.com</a>`,
		},
		{
			Question: "Where is your office?",
			Answer:   "Our team is fully remote!",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}

}
