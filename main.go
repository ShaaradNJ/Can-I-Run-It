package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a game name:")
	game_name, _ := reader.ReadString('\n')

	game_name = strings.TrimSpace(game_name)

	gameRequirements, err := FetchGameRequirements(game_name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	PrintASCIIArtWithInfo()
	fmt.Println("Minimum System Requirements:")
	fmt.Println(gameRequirements.MinCPU)
	fmt.Println(gameRequirements.MinRAM)
	fmt.Println(gameRequirements.MinVideoCard)
	fmt.Println(gameRequirements.MinDedicatedVideoRAM)
	fmt.Println(gameRequirements.MinDiskSpace)
	fmt.Println(gameRequirements.MinOS)
}
