package rockies

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// Header element
type Header struct {
	XMLName         xml.Name `xml:"header"`
	Name            string   `xml:"name,attr"`
	Acknowledgement string   `xml:"acknowledgment,attr"`
	History         string   `xml:"history,attr"`
	Access          string   `xml:"access,attr"`
	Camping         string   `xml:"camping,attr"`
	Rock            string   `xml:"rock,attr"`
	Sun             string   `xml:"sun,attr"`
	Walk            string   `xml:"walk,attr"`
	Intro           string   `xml:"intro,attr"`
}

// Image element
type Image struct {
	ID    string `xml:"id,attr"`
	Src   string `xml:"src,attr"`
	Width string `xml:"width,attr"`
}

// Climb element
type Climb struct {
	ID     string `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	Length string `xml:"length,attr"`
	Stars  string `xml:"stars,attr"`
	Grade  string `xml:"grade,attr"`
	FA     string `xml:"fa,attr"`
	Text   string `xml:",chardata"`
}

// Problem element
type Problem struct {
	ID     string `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	Length string `xml:"length,attr"`
	Stars  string `xml:"stars,attr"`
	Grade  string `xml:"grade,attr"`
	Extra  string `xml:"extra,attr"`
	FA     string `xml:"fa,attr"`
	Text   string `xml:",chardata"`
}

// Text element
type Text struct {
	Class string `xml:"class,attr"`
	Text  string `xml:",chardata"`
}

// Crag element
type Crag struct {
	Xmlid string `xml:"xmlid,attr"`
	Name  string `xml:"name,attr"`
}

// Mixed element
type Mixed struct {
	Type  string
	Value interface{}
}

// Guide element
type Guide struct {
	XMLName  xml.Name `xml:"guide"`
	Contents []Mixed  `xml:",any"`
}

// UnmarshalXML implementation for guidebook XML
func (m *Mixed) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	//fmt.Printf("start.Name.Local=%s\n", start.Name.Local)
	switch start.Name.Local {
	case "header":
		var header Header
		err := d.DecodeElement(&header, &start)
		if err != nil {
			return err
		}
		m.Value = header
		m.Type = start.Name.Local
	case "image":
		var image Image
		err := d.DecodeElement(&image, &start)
		if err != nil {
			return err
		}
		m.Value = image
		m.Type = start.Name.Local
	case "text":
		var text Text
		err := d.DecodeElement(&text, &start)
		if err != nil {
			return err
		}
		m.Value = text
		m.Type = start.Name.Local
	case "climb":
		var climb Climb
		err := d.DecodeElement(&climb, &start)
		if err != nil {
			return err
		}
		m.Value = climb
		m.Type = start.Name.Local
	case "problem":
		var problem Problem
		err := d.DecodeElement(&problem, &start)
		if err != nil {
			return err
		}
		m.Value = problem
		m.Type = start.Name.Local
	case "crag":
		var crag Crag
		err := d.DecodeElement(&crag, &start)
		if err != nil {
			return err
		}
		m.Value = crag
		m.Type = start.Name.Local
	default:
		return fmt.Errorf("unknown element: %s", start)
	}
	return nil
}

// ReadXML reads the Confluence XML from file with path
func ReadXML(path string) (*Guide, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read %s", path)
	}

	var guide Guide
	err = xml.Unmarshal(b, &guide)
	if err != nil {
		return nil, err
	}

	return &guide, nil
}
