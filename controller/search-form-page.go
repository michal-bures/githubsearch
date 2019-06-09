package controller

import "net/http"

func SearchFormPage(w http.ResponseWriter) {
	err := searchPageTemplate.Execute(w, nil)
	handleError(w, err)
}
