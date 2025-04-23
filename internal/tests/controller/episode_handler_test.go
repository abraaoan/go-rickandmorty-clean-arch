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

func TestGetEpisodeHandler_Successfully(t *testing.T) {
	mockUC := &mocks.EpisodeUseCaseMock{
		GetEpisodeFn: func(id int) (*entity.Episode, error) {
			return &entity.Episode{ID: id, Name: "Episode 01"}, nil
		},
	}

	handler := controller.NewEpisodeHandler(mockUC)
	req := httptest.NewRequest(http.MethodGet, "Episode?id=1", nil)
	rr := httptest.NewRecorder()

	handler.GetEpisode(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetEpisodesHandler_Successfully(t *testing.T) {
	mockUC := &mocks.EpisodeUseCaseMock{
		GetEpisodesFn: func() (*entity.EpisodeList, error) {
			return &entity.EpisodeList{
				Results: []entity.Episode{
					{ID: 1, Name: "Episode 01"},
					{ID: 2, Name: "Episode 02"},
				},
			}, nil
		},
	}

	handler := controller.NewEpisodeHandler(mockUC)
	req := httptest.NewRequest(http.MethodGet, "Episodes", nil)
	rr := httptest.NewRecorder()

	handler.GetEpisodes(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
