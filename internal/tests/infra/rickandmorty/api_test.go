package rickandmorty_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"
	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/infra/rickandmorty"
	"github.com/stretchr/testify/assert"
)

func TestGetCharacters_UseCase(t *testing.T) {
	response := entity.CharacterList{
		Results: []entity.Character{
			{ID: 1, Name: "Rick", Status: "Alive", Species: "Human"},
		},
	}
	jsonData, _ := json.Marshal(response)
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}))
	defer server.Close()

	cache := rickandmorty.NewURLCache(1 * time.Second)
	client := rickandmorty.NewApiClient(server.URL, cache)

	list1, err := client.GetCharacters(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, callCount)
	assert.Equal(t, "Rick", list1.Results[0].Name)

	list2, err := client.GetCharacters(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, callCount)
	assert.Equal(t, "Rick", list2.Results[0].Name)
}
