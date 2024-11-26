package server

import (
	"log"
	"net/http"
)

type Server struct {
	Mux    *http.ServeMux
	Addr   string
	Logger *log.Logger
}

func NewServer(addr string, logger *log.Logger) *Server {
	svr := &Server{
		Addr:   addr,
		Logger: logger,
	}
	return svr
}

func (svr *Server) setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", Chain(HealthHandler, Logging(svr.Logger)))
	mux.HandleFunc("GET /films/{name}/watch", Chain(FileStreamer, Logging(svr.Logger)))
	mux.HandleFunc("GET /films/{name}", Chain(InfoHandler, Logging(svr.Logger)))
	mux.HandleFunc("GET /films", Chain(ListMovies, Logging(svr.Logger)))
}

func (svr *Server) Serve() {
	mux := http.NewServeMux()
	svr.setupRoutes(mux)
	log.Println("Starting server on " + svr.Addr)
	log.Fatal(http.ListenAndServe(svr.Addr, mux))
}
