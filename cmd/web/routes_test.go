package main

import (
	"fmt"
	"testing"

	"github.com/amitkrkun/bookings/internal/config"
	"github.com/go-chi/chi/v5"
	// version can impact the test results
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing; test passed
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, type is %T", v))
	}
}

// func TestRoutes(t *testing.T) {
// 	var app config.AppConfig

// 	// app
// 	app = config.AppConfig{}

// 	//routes
// 	mux := routes(&app)

// 	//  *chi.Mux
// 	switch v := mux.(type) {
// 	case *chi.Mux:
// 		// do nothing
// 	default:
// 		t.Errorf("type is not *chi.Mux, but is %T", v)
// 	}
// }
