package main

import "encoding/xml"

func (da *DivArray) getById(id string) Div {

	for _, d := range *da {

		if d.AttrId == id {
			return d
		}

		backDiv := d.Div.getById(id)

		if backDiv.AttrId != "" {
			return backDiv
		}
	}

	return Div{}
}

type HTML struct {
	XMLName xml.Name `xml:"html"`
	Head Head `xml:"head"`
	Body Body `xml:"body"`
}

type Body struct {
	Div DivArray `xml:"div"`
}

type Div struct {
	AttrId string `xml:"id,attr"`
	Div DivArray `xml:"div"`
	H2 []H2 `xml:"h2"`
	Table []Table `xml:"table"`
	Value string`xml:",chardata"`
}

type H2 struct {
	Span []Span `xml:"span"`
}

type Span struct {
	Class string `xml:"class,attr"`
	Id string `xml:"id,attr"`
	Value string `xml:",chardata"`
	A []A `xml:"a"`
}

type DivArray []Div

type Head struct {
	Title string `xml:"title"`
}

type Table struct {
	Body TableBody `xml:"tbody"`
}

type TableBody struct {
	Row []Row `xml:"tr"`
}

type Row struct {
	HeaderData []TableHeader `xml:"th"`
	Data []TableData `xml:"td"`
}

type TableHeader struct {
	Colspan int `xml:"colspan,attr"`
	Rowspan int `xml:"rowspan,attr"`
	Div Div `xml:"div"`
	Value string `xml:",chardata"`
}

type TableData struct {
	A []A `xml:"a"`
	Span []Span `xml:"span"`
	Value string `xml:",chardata"`
}

type A struct {
	Href string `xml:"href,attr"`
	Title string `xml:"title,attr"`
	Value string `xml:",chardata"`
}

