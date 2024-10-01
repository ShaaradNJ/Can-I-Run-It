package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	PrintASCIIArtWithInfo()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a game name:")
	game_name, _ := reader.ReadString('\n')

	game_name = strings.TrimSpace(game_name)

	gameRequirements, err := FetchGameRequirements(game_name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Minimum System Requirements:")
	fmt.Println("CPU:", gameRequirements.MinCPU)
	fmt.Println("RAM:", gameRequirements.MinRAM)
	fmt.Println("Video Card:", gameRequirements.MinVideoCard)
	fmt.Println("Dedicated Video RAM:", gameRequirements.MinDedicatedVideoRAM)
	fmt.Println("Disk Space:", gameRequirements.MinDiskSpace)
	fmt.Println("Operating System:", gameRequirements.MinOS)
}
