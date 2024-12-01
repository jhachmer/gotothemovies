package server

import (
	"log"
	"net/http"
)

// Server struct with Address and Logger fields
type Server struct {
	Addr   string
	Logger *log.Logger
}

// NewServer returns a new Server instance with given Address and Logger values
func NewServer(addr string, logger *log.Logger) *Server {
	svr := &Server{
		Addr:   addr,
		Logger: logger,
	}
	return svr
}

// setupRoutes initializes the URL Routes of the Server
// Handlers are wrapped with Middleware
func (svr *Server) setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", Chain(HealthHandler, Logging(svr.Logger)))
	mux.HandleFunc("GET /films/{name}/watch", Chain(FileStreamer, Logging(svr.Logger)))
	mux.HandleFunc("GET /films/{name}", Chain(InfoHandler, Logging(svr.Logger)))
	mux.HandleFunc("GET /films", Chain(ListMovies, Logging(svr.Logger)))
}

// Serve calls setup functions and spins up the Server
func (svr *Server) Serve() {
	mux := http.NewServeMux()
	svr.setupRoutes(mux)
	svr.Logger.Println("Starting server on " + svr.Addr)
	log.Fatal(http.ListenAndServe(svr.Addr, mux))
}
