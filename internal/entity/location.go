package entity

type Location struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Locationtype string   `json:"type"`
	Dimension    string   `json:"dimension"`
	Residents    []string `json:"residents"`
	Url          string   `json:"url"`
	Created      string   `json:"created"`
}

type LocationList struct {
	Results []Location `json:"results"`
}
