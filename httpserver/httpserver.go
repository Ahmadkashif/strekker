package httpserver

import (
	"net/http"
)

type Server struct {
	http.Server
}

func NewServer(mux *http.ServeMux, addr string) *Server {
	return &Server{
		Server: http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}

func (s *Server) Get(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		handlerFunc(w, r)
	})
}

func (s *Server) Post(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		handlerFunc(w, r)
	})
}

func (s *Server) Put(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		handlerFunc(w, r)
	})
}

func (s *Server) Delete(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		handlerFunc(w, r)
	})
}
