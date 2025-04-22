package usecase

import "github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"

type characterRepository interface {
	GetCharacter(id int) (*entity.Character, error)
	GetCharacters(page int) (*entity.CharacterList, error)
}

type CharacterUseCase struct {
	repo characterRepository
}

func NewCharacterUseCase(r characterRepository) *CharacterUseCase {
	return &CharacterUseCase{repo: r}
}

func (uc *CharacterUseCase) GetCharacter(id int) (*entity.Character, error) {
	return uc.repo.GetCharacter(id)
}

func (uc *CharacterUseCase) GetCharacters(page int) (*entity.CharacterList, error) {
	return uc.repo.GetCharacters(page)
}
