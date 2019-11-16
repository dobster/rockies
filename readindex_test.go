package rockies

import "testing"

func TestPages(t *testing.T) {
	pages, err := Pages("json/index.json")
	if err != nil {
		t.Errorf("unable to load index.json: %v", err)
	}

	if len(pages) != 5 {
		t.Errorf("length of pages != 5")
	}

	page3 := pages[3]
	if page3.Title != "Blue Mountains" {
		t.Errorf("expected title to be Blue Mountains, found %s", page3.Title)
	}

	if page3.ID != "1310731" {
		t.Errorf("expected ID to be 1310731, found %s", page3.ID)
	}
}

func TestPageMaps(t *testing.T) {
	m, err := PageMaps("json/index.json")
	if err != nil {
		t.Errorf("unable to load index.json: %v", err)
	}

	bungers := m["Warrumbungles"]
	if bungers.Title != "Warrumbungles" {
		t.Errorf("expected title to be Warrumbungles, found %s", bungers.Title)
	}

	if _, ok := m["Rubbish"]; ok {
		t.Errorf("was expecting no match for Rubbish")
	}
}
