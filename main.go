package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Racket struct {
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	Price      string `json:"price"`
	ImageUrl   string `json:"imageUrl"`
	RacketPage string `json:"racketPage"`
	Weight     string `json:"weight"`
	Shape      string `json:"shape"`
	Material   string `json:"material"`
	Series     string `json:"series"`
}

var series []string
var brand = "Starvie"
var mainURL = "https://starvie.com/en"

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

func scrapMainPageOfSeries(url string) {
	var rackets []Racket
	racketSeries := ""

	c := colly.NewCollector()
	c.OnHTML("h1.banner__title", func(e *colly.HTMLElement) {
		racketSeries = strings.ReplaceAll(e.Text, " ", "")
		racketSeries = strings.TrimSpace(racketSeries)
	})

	c.OnHTML("div.card", func(e *colly.HTMLElement) {
		item := Racket{
			Brand:      brand,
			Model:      strings.ReplaceAll(e.ChildText("a"), " ", ""),
			Price:      e.ChildText("span.price__regular"),
			ImageUrl:   e.ChildAttr("img", "src"),
			RacketPage: mainURL + e.ChildAttr("a", "href"),
		}
		item.Series = racketSeries
		item.ImageUrl = strings.SplitN(item.ImageUrl, "?", 2)[0]
		item.ImageUrl = strings.TrimPrefix(item.ImageUrl, "//")
		rackets = append(rackets, item)

	})

	if err := c.Visit(url); err != nil {
		fmt.Println("Visit error: ", err)
	}

	content, err := json.Marshal(rackets)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if err := os.WriteFile(brand+"Racket"+racketSeries+".json", content, 0644); err != nil {
		fmt.Println("OS writing error: ", err)
	}

}

func main() {
	importSeries(mainURL)
	/*
		for i := range series {
			scrapMainPageOfSeries("https://starvie.com/en/collections/" + series[i])
		}
	*/
	scrapMainPageOfSeries("https://starvie.com/en/collections/super-pro-line")
}
