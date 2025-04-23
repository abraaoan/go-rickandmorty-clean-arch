package usecase_test

import (
	"errors"
	"testing"

	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"
	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/tests/mocks"
	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetCharacter_Success(t *testing.T) {
	mockRepo := &mocks.CharacterUseCaseMock{
		GetCharacterFn: func(id int) (*entity.Character, error) {
			return &entity.Character{
				ID:      id,
				Name:    "Rick",
				Status:  "Alive",
				Species: "Human",
			}, nil
		},
	}

	uc := usecase.NewCharacterUseCase(mockRepo)
	character, err := uc.GetCharacter(1)

	assert.NoError(t, err)
	assert.Equal(t, "Rick", character.Name)
}

func TestGetCharacter_InvalidID(t *testing.T) {
	mockRepo := &mocks.CharacterUseCaseMock{
		GetCharacterFn: func(id int) (*entity.Character, error) {
			return nil, errors.New("Invalid ID")
		},
	}

	uc := usecase.NewCharacterUseCase(mockRepo)
	character, err := uc.GetCharacter(-1)

	assert.Error(t, err)
	assert.Nil(t, character)
}

func TestGetCharacteres_InvalidPage(t *testing.T) {
	mockRepo := &mocks.CharacterUseCaseMock{
		GetCharactersFn: func(page int) (*entity.CharacterList, error) {
			return nil, errors.New("invalid page")
		},
	}

	uc := usecase.NewCharacterUseCase(mockRepo)
	list, err := uc.GetCharacters(-1)

	assert.Error(t, err)
	assert.Nil(t, list)
}
