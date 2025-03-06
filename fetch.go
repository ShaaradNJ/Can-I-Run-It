package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type GameRequirements struct {
	MinCPU               string
	MinRAM               string
	MinVideoCard         string
	MinDedicatedVideoRAM string
	MinDiskSpace         string
	MinOS                string
	FinalGameName        string
}

func FetchGameRequirements(game_name string) (GameRequirements, error) {
	var gameRequirements GameRequirements
	var visit_here string
	first_char := string(game_name[0])

	domain := "https://www.systemrequirementslab.com"
	c := colly.NewCollector()
	c.SetClient(&http.Client{
		Timeout: 30 * time.Second,
	})

	c.OnHTML("a", func(h *colly.HTMLElement) {
		titleAttr := h.Attr("title")
		gameTitle := h.Text

		if strings.Contains(h.Attr("href"), "/requirements/") &&
			(strings.EqualFold(titleAttr, game_name) || strings.EqualFold(gameTitle, game_name)) {
			visit_here = h.Attr("href")
			fmt.Println("Game found:", gameTitle, "URL:", visit_here)
		}
	})

	c.OnHTML("div.col.col-8 ul li", func(h *colly.HTMLElement) {
		content := strings.TrimSpace(h.Text)
		switch {
		case strings.Contains(content, "CPU") && !strings.Contains(content, "CPU SPEED"):
			gameRequirements.MinCPU = strings.TrimSpace(content)
		case strings.Contains(content, "RAM"):
			gameRequirements.MinRAM = strings.TrimSpace(content)
		case strings.Contains(content, "VIDEO CARD"):
			gameRequirements.MinVideoCard = strings.TrimSpace(content)
		case strings.Contains(content, "DEDICATED VIDEO RAM"):
			gameRequirements.MinDedicatedVideoRAM = strings.TrimSpace(content)
		case strings.Contains(content, "FREE DISK SPACE"):
			gameRequirements.MinDiskSpace = strings.TrimSpace(content)
		case strings.Contains(content, "OS"):
			gameRequirements.MinOS = strings.TrimSpace(content)
		}
	})

	c.OnHTML("h2 em", func(h *colly.HTMLElement) {
		gameRequirements.FinalGameName = strings.TrimSpace(h.Text)
	})

	game_list_url := fmt.Sprintf("https://www.systemrequirementslab.com/all-games-list/?filter=%s", first_char)
	err := c.Visit(game_list_url)
	if err != nil {
		return GameRequirements{}, fmt.Errorf("error visiting the game list page: %v", err)
	}

	if visit_here != "" {
		game_url := domain + visit_here
		fmt.Println("Visiting:", game_url)
		err := c.Visit(game_url)
		if err != nil {
			return GameRequirements{}, fmt.Errorf("error visiting the game details page: %v", err)
		}
	} else {
		return GameRequirements{}, fmt.Errorf("no matching game found")
	}

	return gameRequirements, nil
}


