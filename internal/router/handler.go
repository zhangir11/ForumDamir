package router

import (
	"fmt"
	"net/http"
)

type MainHandler struct {
	routes *Routes
}

func NewMainHandler() (*MainHandler, error) {
	r, err := NewRoutes()
	if err != nil {
		return nil, fmt.Errorf("MainHandler.NewMainHandler: %w", err)
	}

	return &MainHandler{
		routes: r,
	}, nil
}

func (m *MainHandler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	return mux
}
