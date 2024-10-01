package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	var game_name string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a game name:")
	game_name, _ = reader.ReadString('\n')
	game_name_formatted := strings.ReplaceAll(game_name, " ", "+")
	fmt.Println("The game name is:", game_name)
	c := colly.NewCollector()
	c.OnHTML("h1", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})
	defer c.Visit("https://technical.city/en/system-requirements")
	fmt.Println(game_name_formatted)
}
