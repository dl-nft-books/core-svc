package opensea

// Metadata is a struct of a json file according to the OpenSea specification
type Metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	FileURL     string `json:"external_url"`
}
