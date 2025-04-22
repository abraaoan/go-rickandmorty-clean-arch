package entity

type Episode struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	AirDate    string   `json:"air_date"`
	Episode    string   `json:"episode"`
	Characters []string `json:"characters"`
	Url        string   `json:"url"`
	Created    string   `json:"created"`
}

type EpisodeList struct {
	Results []Episode `json:"results"`
}
