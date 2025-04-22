package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"
	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/interface/controller"
	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetCharacterHandler(t *testing.T) {
	mockUC := &mocks.CharacterUseCaseMock{
		GetCharacterFn: func(id int) (*entity.Character, error) {
			return &entity.Character{ID: 1, Name: "Rick"}, nil
		},
	}

	handler := controller.NewCharacterHandler(mockUC)
	req := httptest.NewRequest(http.MethodGet, "/character?id=1", nil)
	rr := httptest.NewRecorder()

	handler.GetCharacter(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetCharacterHandler_WrongParameter(t *testing.T) {
	mockUC := &mocks.CharacterUseCaseMock{
		GetCharacterFn: func(id int) (*entity.Character, error) {
			return nil, errors.New("Invalid ID")
		},
	}

	handler := controller.NewCharacterHandler(mockUC)
	req := httptest.NewRequest(http.MethodGet, "/character?id=ABC", nil)
	rr := httptest.NewRecorder()

	handler.GetCharacter(rr, req)
	body := rr.Body.String()

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, body, "Invalid ID")
}
