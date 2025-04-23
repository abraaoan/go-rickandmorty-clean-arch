package usecase_test

import (
	"errors"
	"testing"

	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"
	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/tests/mocks"
	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetLocation_Success(t *testing.T) {
	mockRepo := &mocks.LocationUseCaseMock{
		GetLocationFn: func(id int) (*entity.Location, error) {
			return &entity.Location{
				ID:   id,
				Name: "Earth",
			}, nil
		},
	}

	uc := usecase.NewLocationUseCase(mockRepo)
	location, err := uc.GetLocation(1)

	assert.NoError(t, err)
	assert.Equal(t, "Earth", location.Name)
}

func TestGetLocation_Invalid(t *testing.T) {
	mockRepo := &mocks.LocationUseCaseMock{
		GetLocationFn: func(id int) (*entity.Location, error) {
			return nil, errors.New("Invalid id")
		},
	}

	uc := usecase.NewLocationUseCase(mockRepo)
	location, err := uc.GetLocation(-1)

	assert.Error(t, err)
	assert.Nil(t, location)
}

func TestGetLocations_Success(t *testing.T) {
	mockRepo := &mocks.LocationUseCaseMock{
		GetLocationsFn: func() (*entity.LocationList, error) {
			return &entity.LocationList{
				Results: []entity.Location{
					{ID: 1, Name: "Earth"},
					{ID: 2, Name: "Mars"},
				},
			}, nil
		},
	}

	uc := usecase.NewLocationUseCase(mockRepo)
	locations, err := uc.GetLocations()

	assert.NoError(t, err)
	assert.Equal(t, len(locations.Results), 2)
	assert.Equal(t, locations.Results[0].Name, "Earth")
	assert.Equal(t, locations.Results[1].Name, "Mars")
}

func TestGetLocations_Invalid(t *testing.T) {
	mockRepo := &mocks.LocationUseCaseMock{
		GetLocationsFn: func() (*entity.LocationList, error) {
			return nil, errors.New("Not found")
		},
	}

	uc := usecase.NewLocationUseCase(mockRepo)
	locations, err := uc.GetLocations()

	assert.Nil(t, locations)
	assert.Error(t, err)
}
