package models

type Metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	FileURL     string `json:"external_url"`
}
