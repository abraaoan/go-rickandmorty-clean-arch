package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"
)

type LocationUseCase interface {
	GetLocation(id int) (*entity.Location, error)
	GetLocations() (*entity.LocationList, error)
}

type locationHandler struct {
	usecase LocationUseCase
}

func NewLocationHandler(uc LocationUseCase) *locationHandler {
	return &locationHandler{usecase: uc}
}

func (h *locationHandler) GetLocation(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	idStr := query.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	location, err := h.usecase.GetLocation(id)
	if err != nil {
		http.Error(w, "Location not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(location)
}

func (h *locationHandler) GetLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := h.usecase.GetLocations()
	if err != nil {
		http.Error(w, "Locations not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(locations)
}
