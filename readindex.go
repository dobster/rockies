package rockies

import (
	"encoding/json"
	"io/ioutil"
)

// Page is page summary
type Page struct {
	Title string `json:"title"`
	ID    string `json:"id"`
}

// Pages loads a JSON array of Page structs
func Pages(name string) ([]Page, error) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var pages []Page
	err = json.Unmarshal(b, &pages)
	if err != nil {
		return nil, err
	}

	return pages, nil
}

// PageMaps loads a JSON array of Page structs and maps them by
func PageMaps(name string) (map[string]Page, error) {
	a, err := Pages(name)
	if err != nil {
		return nil, err
	}

	pages := make(map[string]Page, len(a))
	for _, page := range a {
		pages[page.Title] = page
	}

	return pages, nil
}
