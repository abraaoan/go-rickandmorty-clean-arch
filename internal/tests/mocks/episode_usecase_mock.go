package mocks

import "github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"

type EpisodeUseCaseMock struct {
	GetEpisodeFn  func(id int) (*entity.Episode, error)
	GetEpisodesFn func() (*entity.EpisodeList, error)
}

func (m *EpisodeUseCaseMock) GetEpisode(id int) (*entity.Episode, error) {
	return m.GetEpisodeFn(id)
}

func (m *EpisodeUseCaseMock) GetEpisodes() (*entity.EpisodeList, error) {
	return m.GetEpisodesFn()
}
