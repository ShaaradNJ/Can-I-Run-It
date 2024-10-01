// fetch.go
package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type GameRequirements struct {
	MinCPU               string
	MinRAM               string
	MinVideoCard         string
	MinDedicatedVideoRAM string
	MinDiskSpace         string
	MinOS                string
}

func FetchGameRequirements(game_name string) (GameRequirements, error) {
	var gameRequirements GameRequirements
	var visit_here string

	domain := "https://www.systemrequirementslab.com"
	c := colly.NewCollector()

	c.OnHTML("a", func(h *colly.HTMLElement) {
		titleAttr := h.Attr("title")
		gameTitle := h.Text

		if strings.Contains(strings.ToLower(titleAttr), strings.ToLower(game_name)) ||
			strings.Contains(strings.ToLower(gameTitle), strings.ToLower(game_name)) {
			visit_here = h.Attr("href")
			fmt.Println("Game found:", gameTitle)
		}
	})

	// Scraping system requirements
	c.OnHTML("li", func(h *colly.HTMLElement) {
		if strings.Contains(h.ChildText("strong"), "CPU") {
			gameRequirements.MinCPU = h.Text
		}
	})
	c.OnHTML("li", func(h *colly.HTMLElement) {
		if strings.Contains(h.ChildText("strong"), "RAM") {
			gameRequirements.MinRAM = h.Text
		}
	})
	c.OnHTML("li", func(h *colly.HTMLElement) {
		if strings.Contains(h.ChildText("strong"), "VIDEO CARD") {
			gameRequirements.MinVideoCard = h.Text
		}
	})
	c.OnHTML("li", func(h *colly.HTMLElement) {
		if strings.Contains(h.ChildText("strong"), "DEDICATED VIDEO RAM") {
			gameRequirements.MinDedicatedVideoRAM = h.Text
		}
	})
	c.OnHTML("li", func(h *colly.HTMLElement) {
		if strings.Contains(h.ChildText("strong"), "FREE DISK SPACE") {
			gameRequirements.MinDiskSpace = h.Text
		}
	})
	c.OnHTML("li", func(h *colly.HTMLElement) {
		if strings.Contains(h.ChildText("strong"), "OS") {
			gameRequirements.MinOS = h.Text
		}
	})

	err := c.Visit("https://www.systemrequirementslab.com/all-games-list")
	if err != nil {
		return GameRequirements{}, fmt.Errorf("error visiting the game list page: %v", err)
	}

	// If a matching game was found
	if visit_here != "" {
		game_url := domain + visit_here
		fmt.Printf("Full game URL: %s\n", game_url)

		err := c.Visit(game_url)
		if err != nil {
			return GameRequirements{}, fmt.Errorf("error visiting the game page: %v", err)
		}
	} else {
		return GameRequirements{}, fmt.Errorf("no matching game found")
	}

	return gameRequirements, nil
}
