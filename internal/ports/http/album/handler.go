package album

import (
	"encoding/json"
	"fmt"
	"github.com/ydimitriou/go-clean-architecture-example/internal/app"
	"github.com/ydimitriou/go-clean-architecture-example/internal/app/album/commands"
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
		fmt.Fprint(w, err.Error())
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
		fmt.Fprint(w, err.Error())
	}
	w.WriteHeader(http.StatusCreated)
}
