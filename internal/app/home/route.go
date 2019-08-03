package home

import (
	"net/http"

	"github.com/pythinh/go-news/internal/app/types"
)

const (
	get = http.MethodGet
)

// NewRouter append new router
func NewRouter(r *[]types.Route) {
	routes := []types.Route{
		{
			Path:    "/",
			Method:  get,
			Handler: newRouter().indexHandler,
		},
		{
			Path:    "/about",
			Method:  get,
			Handler: newRouter().aboutHandler,
		},
	}

	*r = append(*r, routes...)
}
