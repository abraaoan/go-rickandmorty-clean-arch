package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"
	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/interface/controller"
	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetLocationHandler(t *testing.T) {
	mockUC := &mocks.LocationUseCaseMock{
		GetLocationFn: func(id int) (*entity.Location, error) {
			return &entity.Location{ID: 1, Name: "Earth"}, nil
		},
	}

	handler := controller.NewLocationHandler(mockUC)
	req := httptest.NewRequest(http.MethodGet, "/location?id=1", nil)
	rr := httptest.NewRecorder()

	handler.GetLocation(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
