package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type GameStats struct {
	GameName    string `json:"game_name"`
	DoorCode    string `json:"door_code"`
	Category    string `json:"category"`
	LaunchCount int    `json:"launch_count"`
}

type Top10Stats struct {
	Period string      `json:"period"`
	Games  []GameStats `json:"games"`
}

func main() {
	// Get the current month name
	currentMonth := time.Now().Format("January")

	// Fetch the Top 10 games for the current month
	url := fmt.Sprintf("https://goldminedoors.com/top10?period=%s", strings.ToLower(currentMonth))
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching top 10 games:", err)
		return
	}
	defer response.Body.Close()

	var top10 Top10Stats
	if err := json.NewDecoder(response.Body).Decode(&top10); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Prepare ANSI art
	ansiArt := generateAnsiArt(top10.Games)

	// Write to top10.ans
	if err := ioutil.WriteFile("top10.ans", []byte(ansiArt), 0644); err != nil {
		fmt.Println("Error writing to top10.ans:", err)
		return
	}

	fmt.Println("Top 10 ANSI art generated successfully.")
}

func generateAnsiArt(games []GameStats) string {
	var sb strings.Builder

	// Start with a header
	sb.WriteString("\x1b[1;37m") // White bold text
	sb.WriteString("╔══════════════════════════════════════╗\n")
	sb.WriteString("║\x1b[1;34m        Top 10 Games for the Month       \x1b[1;37m║\n")
	sb.WriteString("╚══════════════════════════════════════╝\n")

	// Add the top 10 games
	for i, game := range games {
		sb.WriteString(fmt.Sprintf("\x1b[1;33m%2d. \x1b[1;32m%-20s\x1b[1;36m %5s \x1b[1;31m(%d)\n",
			i+1, game.GameName, game.DoorCode, game.LaunchCount))
	}

	// Footer
	sb.WriteString("\x1b[1;37m╚══════════════════════════════════════╝\n")
	sb.WriteString("\x1b[0m") // Reset colors

	return sb.String()
}
