package usecase

import "github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"

type epsiodeRepository interface {
	GetEpisode(id int) (*entity.Episode, error)
	GetEpisodes() (*entity.EpisodeList, error)
}

type EpisodeUseCase struct {
	repo epsiodeRepository
}

func NewEpisodeUseCase(r epsiodeRepository) *EpisodeUseCase {
	return &EpisodeUseCase{repo: r}
}

func (uc EpisodeUseCase) GetEpisode(id int) (*entity.Episode, error) {
	return uc.repo.GetEpisode(id)
}

func (uc EpisodeUseCase) GetEpisodes() (*entity.EpisodeList, error) {
	return uc.repo.GetEpisodes()
}
