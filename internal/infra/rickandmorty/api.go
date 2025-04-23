package rickandmorty

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/entity"
)

type APIClient struct {
	baseURL string
	client  *http.Client
	cache   *URLCache
}

func NewApiClient(baseURL string, cache *URLCache) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		client:  &http.Client{},
		cache:   cache,
	}
}

func (api *APIClient) getJsonWithCache(url string, dest any) error {
	if data, ok := api.cache.Get(url); ok {
		return json.Unmarshal(data, dest)
	}

	resp, err := api.client.Get(url)
	if err != nil {
		return fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}

	api.cache.Set(url, body)
	return json.Unmarshal(body, dest)
}

func (api *APIClient) GetCharacter(id int) (*entity.Character, error) {
	url := fmt.Sprintf("%s/character/%d", api.baseURL, id)

	var character entity.Character
	err := api.getJsonWithCache(url, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

func (api *APIClient) GetCharacters(page int) (*entity.CharacterList, error) {
	url := fmt.Sprintf("%s/character?page=%d", api.baseURL, page)
	var characters entity.CharacterList
	err := api.getJsonWithCache(url, &characters)
	if err != nil {
		return nil, err
	}

	return &characters, nil
}

func (api *APIClient) GetLocation(id int) (*entity.Location, error) {
	url := fmt.Sprintf("%s/location/%d", api.baseURL, id)
	var location entity.Location
	err := api.getJsonWithCache(url, &location)
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func (api *APIClient) GetLocations() (*entity.LocationList, error) {
	url := fmt.Sprintf("%s/location", api.baseURL)
	var locations entity.LocationList
	err := api.getJsonWithCache(url, &locations)
	if err != nil {
		return nil, err
	}
	return &locations, nil
}

func (api *APIClient) GetEpisode(id int) (*entity.Episode, error) {
	url := fmt.Sprintf("%s/episode/%d", api.baseURL, id)
	var episode entity.Episode
	err := api.getJsonWithCache(url, &episode)
	if err != nil {
		return nil, err
	}
	return &episode, nil
}

func (api *APIClient) GetEpisodes() (*entity.EpisodeList, error) {
	url := fmt.Sprintf("%s/episode", api.baseURL)

	var episodes entity.EpisodeList
	err := api.getJsonWithCache(url, &episodes)
	if err != nil {
		return nil, err
	}

	return &episodes, nil
}
