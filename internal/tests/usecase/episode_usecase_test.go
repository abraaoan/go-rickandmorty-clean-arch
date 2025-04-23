package usecase_test

import (
	"errors"
	"testing"

	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"
	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/tests/mocks"
	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetEpisode_Successfully(t *testing.T) {
	mockRepo := &mocks.EpisodeUseCaseMock{
		GetEpisodeFn: func(id int) (*entity.Episode, error) {
			return &entity.Episode{ID: id, Name: "Episode 01"}, nil
		},
	}

	uc := usecase.NewEpisodeUseCase(mockRepo)
	episode, err := uc.GetEpisode(1)

	assert.NoError(t, err)
	assert.Equal(t, episode.ID, 1)
	assert.Equal(t, episode.Name, "Episode 01")
}

func TestGetEpisode_Invalid(t *testing.T) {
	mockRepo := &mocks.EpisodeUseCaseMock{
		GetEpisodeFn: func(id int) (*entity.Episode, error) {
			return nil, errors.New("Invalid ID")
		},
	}

	uc := usecase.NewEpisodeUseCase(mockRepo)
	episode, err := uc.GetEpisode(-1)

	assert.Error(t, err)
	assert.Nil(t, episode)
}

func TestGetEpisodes_Successfully(t *testing.T) {
	mockRepo := &mocks.EpisodeUseCaseMock{
		GetEpisodesFn: func() (*entity.EpisodeList, error) {
			return &entity.EpisodeList{
				Results: []entity.Episode{
					{ID: 1, Name: "Episode 01"},
					{ID: 2, Name: "Episode 02"},
				},
			}, nil
		},
	}

	uc := usecase.NewEpisodeUseCase(mockRepo)
	episodes, err := uc.GetEpisodes()

	assert.NoError(t, err)
	assert.Equal(t, len(episodes.Results), 2)
	assert.Equal(t, episodes.Results[0].Name, "Episode 01")
	assert.Equal(t, episodes.Results[1].Name, "Episode 02")
	assert.Equal(t, episodes.Results[0].ID, 1)
	assert.Equal(t, episodes.Results[1].ID, 2)
}
