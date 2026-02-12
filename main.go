package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

var series []string

func importSeries(url string) {
	c := colly.NewCollector()

	c.OnHTML("ul.dropdown__nav a.reversed-link", func(e *colly.HTMLElement) {
		title := e.Text
		title = strings.ToLower(title)
		if !strings.Contains(title, "line") {
			return
		}
		title = strings.ReplaceAll(title, " ", "-")
		series = append(series, title)
	})

	if err := c.Visit(url); err != nil {
		fmt.Println("Visit error: ", err)
	}
}

func main() {
	importSeries("https://starvie.com/en")
	/*
		for i := range series {
			fmt.Println("https://starvie.com/en/collections/" + series[i])
		}

	*/
}
