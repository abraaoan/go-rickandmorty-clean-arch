package rickandmorty

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"
)

type APIClient struct {
	baseURL string
	client  *http.Client
}

func NewApiClient() *APIClient {
	return &APIClient{
		baseURL: "https://rickandmortyapi.com/api",
		client:  &http.Client{},
	}
}

func (api *APIClient) GetCharacter(id int) (*entity.Character, error) {
	url := fmt.Sprintf("%s/character/%d", api.baseURL, id)
	resp, err := api.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var character entity.Character
	if err := json.NewDecoder(resp.Body).Decode(&character); err != nil {
		return nil, err
	}

	return &character, nil
}

func (api *APIClient) GetCharacters(page int) (*entity.CharacterList, error) {
	url := fmt.Sprintf("%s/character?page=%d", api.baseURL, page)
	resp, err := api.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var characters entity.CharacterList
	if err := json.NewDecoder(resp.Body).Decode(&characters); err != nil {
		return nil, err
	}

	return &characters, nil
}

func (api *APIClient) GetLocation(id int) (*entity.Location, error) {
	url := fmt.Sprintf("%s/location/%d", api.baseURL, id)
	resp, err := api.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var location entity.Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}
	return &location, nil
}

func (api *APIClient) GetLocations() (*entity.LocationList, error) {
	url := fmt.Sprintf("%s/location", api.baseURL)
	resp, err := api.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var locations entity.LocationList
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		return nil, err
	}
	return &locations, nil
}

func (api *APIClient) GetEpisode(id int) (*entity.Episode, error) {
	url := fmt.Sprintf("%s/episode/%d", api.baseURL, id)
	resp, err := api.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var episode entity.Episode
	if err := json.NewDecoder(resp.Body).Decode(&episode); err != nil {
		return nil, err
	}
	return &episode, nil
}

func (api *APIClient) GetEpisodes() (*entity.EpisodeList, error) {
	url := fmt.Sprintf("%s/episode", api.baseURL)
	resp, err := api.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var episodes entity.EpisodeList
	if err := json.NewDecoder(resp.Body).Decode(&episodes); err != nil {
		return nil, err
	}

	return &episodes, nil
}
