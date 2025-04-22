package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"
)

type EpisodeUsecase interface {
	GetEpisode(id int) (*entity.Episode, error)
	GetEpisodes() (*entity.EpisodeList, error)
}

type EpisodeHandler struct {
	usecase EpisodeUsecase
}

func NewEpisodeHandler(uc EpisodeUsecase) *EpisodeHandler {
	return &EpisodeHandler{usecase: uc}
}

func (h *EpisodeHandler) GetEpisode(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	idStr := query.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	epsisode, err := h.usecase.GetEpisode(id)
	if err != nil {
		http.Error(w, "Episode not found", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(epsisode)
}

func (h *EpisodeHandler) GetEpisodes(w http.ResponseWriter, r *http.Request) {
	episodes, err := h.usecase.GetEpisodes()
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(episodes)
}
