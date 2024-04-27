package specs

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"std_go_boilerplate/internal/router"
	"testing"
)

func TestGroup(t *testing.T) {
	mainRouter := http.NewServeMux()

	prefix := "/api"
	handler := func(r *http.ServeMux) {
		r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Test endpoint")
		})
	}

	middleware1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(401)
			next.ServeHTTP(w, r)
		})
	}

	subRouter := http.NewServeMux()
	handler(subRouter)

	router.Group(mainRouter, prefix, handler, middleware1)

	req, err := http.NewRequest("GET", "/api/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	mainRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != 401 {
		t.Errorf("handler returned wrong status code: got %v want %v", status, 401)
	}

	expected := "Test endpoint"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGroupChain(t *testing.T) {
	mainRouter := http.NewServeMux()

	prefix := "/api"
	handler := func(r *http.ServeMux) {
		r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Test endpoint")
		})
	}

	middleware1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(401)
			next.ServeHTTP(w, r)
		})
	}
	middleware2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			next.ServeHTTP(w, r)
		})
	}
	subRouter := http.NewServeMux()
	handler(subRouter)

	router.Group(mainRouter, prefix, handler, middleware1, middleware2)

	req, err := http.NewRequest("GET", "/api/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	mainRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != 500 {
		t.Errorf("handler returned wrong status code: got %v want %v", status, 500)
	}

	expected := "Test endpoint"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
