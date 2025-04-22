package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"
	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/interface/controller/others"
)

type CharacterUseCase interface {
	GetCharacter(id int) (*entity.Character, error)
	GetCharacters(page int) (*entity.CharacterList, error)
}

type CharacterHandler struct {
	usecase CharacterUseCase
}

func NewCharacterHandler(uc CharacterUseCase) *CharacterHandler {
	return &CharacterHandler{usecase: uc}
}

func (h *CharacterHandler) GetCharacter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	idstr := query.Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	character, err := h.usecase.GetCharacter(id)
	if err != nil {
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(character)
}

func (h *CharacterHandler) GetCharacters(w http.ResponseWriter, r *http.Request) {
	page, err := others.ParsePageParam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	characters, err := h.usecase.GetCharacters(page)
	if err != nil {
		http.Error(w, "Failed to fetch characters", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(characters)
}
