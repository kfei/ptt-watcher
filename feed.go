package main

import (
	"bytes"
	"encoding/xml"
	"log"
)

type Atom1 struct {
	XMLName   xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
	Title     string   `xml:"title"`
	Subtitle  string   `xml:"subtitle"`
	Id        string   `xml:"id"`
	Updated   string   `xml:"updated"`
	Rights    string   `xml:"rights"`
	Link      Link     `xml:"link"`
	Author    Author   `xml:"author"`
	EntryList []Entry  `xml:"entry"`
}

type Link struct {
	Href string `xml:"href,attr"`
}

type Author struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

type Entry struct {
	Title     string `xml:"title"`
	Summary   string `xml:"summary"`
	Content   string `xml:"content"`
	Id        string `xml:"id"`
	Published string `xml:"published"`
	Updated   string `xml:"updated"`
	Link      Link   `xml:"link"`
	Author    Author `xml:"author"`
}

func parseAtom(content []byte) (Atom1, error) {
	// Contents from Ptt will sometimes contain the escape character (\x1b),
	// which will leads xml.Unmarshal to fail, so here we strip them first.
	safeContent := bytes.Replace(content, []byte("\x1b"), []byte(""), -1)

	a := Atom1{}
	err := xml.Unmarshal(safeContent, &a)
	if err != nil {
		log.Println(err)
		return Atom1{}, err
	}
	return a, nil
}
