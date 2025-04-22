package entity

type Character struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
}

type CharacterList struct {
	Results []Character `json:"results"`
}
