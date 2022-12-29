package router

import (
	"net/http"
	"text/template"
)

type Routes struct {
	templates *template.Template
}

func NewRoutes() (*Routes, error) {
	r := &Routes{}
	return r, nil
}

func (r *Routes) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", nil)

	mux.HandleFunc("/signup", nil)

	mux.HandleFunc("/signin", nil)

	mux.HandleFunc("/newPost", nil)
	mux.HandleFunc("/userPage", nil)
	mux.HandleFunc("/post", nil)
}
