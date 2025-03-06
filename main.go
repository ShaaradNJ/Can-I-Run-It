package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ConnectDB() // Ensure MongoDB is connected

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a game name:")
	game_name, _ := reader.ReadString('\n')
	game_name = strings.TrimSpace(game_name)

	gameRequirements, err := FetchGameRequirements(game_name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Save the game requirements to MongoDB
	err = SaveGameRequirements(gameRequirements)
	if err != nil {
		fmt.Println("Error saving to MongoDB:", err)
	} else {
		fmt.Println("Game requirements saved successfully")
	}

	PrintASCIIArtWithInfo()
	fmt.Println("Minimum System Requirements for:", gameRequirements.FinalGameName)
	fmt.Println(gameRequirements.MinCPU)
	fmt.Println(gameRequirements.MinRAM)
	fmt.Println(gameRequirements.MinVideoCard)
	fmt.Println(gameRequirements.MinDedicatedVideoRAM)
	fmt.Println(gameRequirements.MinDiskSpace)
	fmt.Println(gameRequirements.MinOS)
	fmt.Println("			**********			")
	fmt.Println()
}
