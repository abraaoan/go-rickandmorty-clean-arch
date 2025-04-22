package usecase

import "github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"

type locationRepository interface {
	GetLocation(id int) (*entity.Location, error)
	GetLocations() (*entity.LocationList, error)
}

type LocationUseCase struct {
	repo locationRepository
}

func NewLocationUseCase(r locationRepository) *LocationUseCase {
	return &LocationUseCase{repo: r}
}

func (uc *LocationUseCase) GetLocation(id int) (*entity.Location, error) {
	return uc.repo.GetLocation(id)
}

func (uc *LocationUseCase) GetLocations() (*entity.LocationList, error) {
	return uc.repo.GetLocations()
}
