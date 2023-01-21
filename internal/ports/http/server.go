package http

import (
	"github.com/gorilla/mux"
	"github.com/ydimitriou/go-clean-architecture-example/internal/app"
	"github.com/ydimitriou/go-clean-architecture-example/internal/ports/http/album"
	"log"
	"net/http"
)

const albumsHTTPPATH = "/albums"

// Server is the http server
type Server struct {
	appServices app.Services
	router      *mux.Router
}

// NewServer HTTP Server constructor
func NewServer(as app.Services) Server {
	router := mux.NewRouter()
	httpServer := Server{
		appServices: as,
		router:      router,
	}
	httpServer.createHTTPRoutes()
	http.Handle("/", httpServer.router)

	return httpServer
}

// createHTTPRoutes generates routes
func (httpServer *Server) createHTTPRoutes() {
	//Commands
	httpServer.router.HandleFunc(albumsHTTPPATH, album.NewHandler(httpServer.appServices.AlbumServices).Create).Methods("POST")
	httpServer.router.HandleFunc(albumsHTTPPATH+"/{"+album.UpdateAlbumID+"}", album.NewHandler(httpServer.appServices.AlbumServices).Update).Methods("PUT")
	httpServer.router.HandleFunc(albumsHTTPPATH+"/{"+album.DeleteAlbumID+"}", album.NewHandler(httpServer.appServices.AlbumServices).Delete).Methods("DELETE")
	// Queries
	httpServer.router.HandleFunc(albumsHTTPPATH, album.NewHandler(httpServer.appServices.AlbumServices).GetAll).Methods("GET")
	httpServer.router.HandleFunc(albumsHTTPPATH+"/{"+album.GetByIDURLParam+"}", album.NewHandler(httpServer.appServices.AlbumServices).GetByID).Methods("GET")
}

// ListenAndServe wraps HTTP listenAndServe (initiate listening for request)
func (httpServer *Server) ListenAndServe(port string) {
	log.Fatal(http.ListenAndServe(port, nil))
}
