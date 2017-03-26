package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/fatih/color"
)

var bold = color.New(color.Bold).SprintFunc()

func fetchPttFeed(url string) (Atom1, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error while fetching feed:", url, "\n", err)
		return Atom1{}, err
	}
	defer response.Body.Close()

	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		log.Println("Error while dumping feed:", url, "\n", err)
		return Atom1{}, err
	}

	feed, err := parseAtom(dump)
	if err != nil {
		log.Println("Error while parsing feed")
	}
	return feed, nil
}

func parsePttTime(timeStr string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Taipei")
	var layout = "2006-01-02T15:04:05-07:00"

	t, err := time.ParseInLocation(layout, strings.TrimSuffix(timeStr, "Z"), loc)
	if err != nil {
		log.Println("Can't parse updated time", err)
		return time.Time{}, err
	}

	return t, nil
}

func filteredAny(str string, filters []Filter) bool {
	if len(filters) == 0 {
		return true
	}

	for _, filter := range filters {
		// Default to true (for filter == "")
		matched := true
		keywords := strings.Fields(string(filter))
		for _, kw := range keywords {
			match := strings.Contains(strings.ToLower(str), strings.ToLower(kw))
			if !match {
				matched = false
				break
			}
		}
		if matched {
			return true
		}
	}

	return false
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
