package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/wahyudin/football-goals-calculator/internal/service"
)

func main() {
	// Validasi command line arguments
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go <teamName> <year>")
	}

	teamName := os.Args[1]
	year, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Invalid year format:", err)
	}

	// Validasi input
	if year < 1900 || year > 2024 {
		log.Fatal("Year must be between 1900 and 2024")
	}

	if teamName == "" {
		log.Fatal("Team name cannot be empty")
	}

	// Inisialisasi service
	footballService := service.NewFootballService()

	// Hitung total gol
	totalGoals, err := footballService.CalculateTotalGoals(teamName, year)
	if err != nil {
		log.Fatal("Error calculating goals:", err)
	}

	fmt.Printf("Total goals scored by %s in %d: %d\n", teamName, year, totalGoals)
}
