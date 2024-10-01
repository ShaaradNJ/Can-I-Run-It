package main

import (
	_ "bufio"
	"fmt"
	_ "os"
	_ "strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	// var game_name string
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter a game name:")
	// game_name, _ = reader.ReadString('\n')
	// game_name_formatted := strings.ReplaceAll(game_name, " ", "+")
	// fmt.Println("The game name is:", game_name)
	c := colly.NewCollector(colly.AllowedDomains(""))
	c.OnHTML("h1", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})
	defer c.Visit("https://www.systemrequirementslab.com/all-games-list")
	// fmt.Println(game_name_formatted)
}
