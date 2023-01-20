package album

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/ydimitriou/go-clean-architecture-example/internal/app"
	"github.com/ydimitriou/go-clean-architecture-example/internal/app/album/commands"
	"github.com/ydimitriou/go-clean-architecture-example/internal/app/album/queries"
	"net/http"
)

// Handler is an album http request handler
type Handler struct {
	appServices app.AlbumServices
}

// NewHandler returns an http request handler
func NewHandler(as app.AlbumServices) Handler {
	return Handler{appServices: as}
}

// CreateAlbumRequest represents the expected model for create requests
type CreateAlbumRequest struct {
	Title       string  `json:"title"`
	Artist      string  `json:"artist"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// Create adds the given album in storage
func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateAlbumRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	album := commands.CreateAlbumRequest{
		Title:       req.Title,
		Artist:      req.Artist,
		Price:       req.Price,
		Description: req.Description,
	}
	err = h.appServices.Commands.CreateAlbumHandler.Handle(album)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}
	w.WriteHeader(http.StatusCreated)
}

// GetAll return all available albums
func (h Handler) GetAll(w http.ResponseWriter, _ *http.Request) {
	albums, err := h.appServices.Queries.GetAllAlbumsHAndler.Handle()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(albums)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
}

const GetByIDURLParam = "albumID"

// GetByID returns the requested by ID album
func (h Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	albumID := v[GetByIDURLParam]
	album, err := h.appServices.Queries.GetAlbumHandler.Handle(queries.GetAlbumRequest{ID: uuid.MustParse(albumID)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	if album == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Album Not Found")
		return
	}
	err = json.NewEncoder(w).Encode(album)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
}
