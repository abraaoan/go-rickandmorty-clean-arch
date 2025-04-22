package mocks

import "github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"

type LocationUseCaseMock struct {
	GetLocationFn  func(id int) (*entity.Location, error)
	GetLocationsFn func() (*entity.LocationList, error)
}

func (m *LocationUseCaseMock) GetLocation(id int) (*entity.Location, error) {
	return m.GetLocationFn(id)
}

func (m *LocationUseCaseMock) GetLocations() (*entity.LocationList, error) {
	return m.GetLocationsFn()
}
