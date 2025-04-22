package mocks

import "github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"

type CharacterUseCaseMock struct {
	GetCharacterFn  func(id int) (*entity.Character, error)
	GetCharactersFn func(page int) (*entity.CharacterList, error)
}

func (m *CharacterUseCaseMock) GetCharacter(id int) (*entity.Character, error) {
	return m.GetCharacterFn(id)
}

func (m *CharacterUseCaseMock) GetCharacters(page int) (*entity.CharacterList, error) {
	return m.GetCharactersFn(page)
}
