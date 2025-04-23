package main

import (
	"log"
	"net/http"
	"time"

	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/infra/rickandmorty"
	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/interface/controller"
	"github.com/abraaoan/go-rickandmorty-clean-arch/internal/usecase"
)

func main() {
	baseUrl := "https://rickandmortyapi.com/api"
	cache := rickandmorty.NewURLCache(5 * time.Minute)
	repo := rickandmorty.NewApiClient(baseUrl, cache)

	// Characters
	characterUseCase := usecase.NewCharacterUseCase(repo)
	characterHandler := controller.NewCharacterHandler(characterUseCase)

	// Location
	locationUseCase := usecase.NewLocationUseCase(repo)
	locationHandler := controller.NewLocationHandler(locationUseCase)

	// Episodes
	episodeUseCase := usecase.NewEpisodeUseCase(repo)
	episodeHandler := controller.NewEpisodeHandler(episodeUseCase)

	// Routes
	http.HandleFunc("/character", characterHandler.GetCharacter)
	http.HandleFunc("/characters", characterHandler.GetCharacters)
	http.HandleFunc("/location", locationHandler.GetLocation)
	http.HandleFunc("/locations", locationHandler.GetLocations)
	http.HandleFunc("/episode", episodeHandler.GetEpisode)
	http.HandleFunc("/episodes", episodeHandler.GetEpisodes)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
