package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {

	var min_cpu string
	var min_ram string
	var min_video_card string
	// var min_dedicated_video_card string
	// var min_pixel_shaders string
	// var min_os string
	// var min_disk_space string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a game name:")
	game_name, _ := reader.ReadString('\n')

	game_name = strings.TrimSpace(game_name)

	// Set the domain
	domain := "https://www.systemrequirementslab.com"
	var visit_here string // This will store the URL for the game

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

	c.OnHTML("li", func(h *colly.HTMLElement) {
		if strings.Contains(h.ChildText("strong"), "CPU") {

			min_cpu = h.Text
		}
	})
	c.OnHTML("li", func(h *colly.HTMLElement) {
		if strings.Contains(h.ChildText("strong"), "RAM") {

			min_ram = h.Text
		}
	})
	c.OnHTML("li", func(h *colly.HTMLElement) {
		if strings.Contains(h.ChildText("strong"), "VIDEO CARD") {

			min_video_card = h.Text
		}
	})

	err := c.Visit("https://www.systemrequirementslab.com/all-games-list")
	if err != nil {
		fmt.Println("Error visiting the page:", err)
	}

	// After visiting the game list, check if a matching game was found
	if visit_here != "" {
		// Concatenate domain and visit_here to form the complete URL
		game_url := domain + visit_here
		fmt.Printf("Full game URL: %s\n", game_url)

		// Visit the game-specific page
		err := c.Visit(game_url)
		if err != nil {
			fmt.Println("Error visiting the game page:", err)
		}
	} else {
		fmt.Println("No matching game found.")
	}
	fmt.Println("")
	fmt.Println(min_cpu)
	fmt.Println(min_ram)
	fmt.Println(min_video_card)
}
