package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Cube struct {
	color    string
	quantity int
	limit    int
}

type Game struct {
	id int
}

func (c *Cube) New(color string, quantity int, limit int) *Cube {
	return &Cube{
		color:    color,
		quantity: quantity,
		limit:    limit,
	}
}

func readFileAsString() *string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	contentAsString := string(content)
	return &contentAsString
}

func parseGame(input *string) {
	gameData := strings.Split(*input, ":")
	// game := strings.Split(gameData[0], " ")[1]
	sets := gameData[1:]
	fmt.Println(sets)
	// fmt.Println(game)
}

func main() {
	file := readFileAsString()
	for _, game := range strings.Split(strings.ReplaceAll(*file, "\r\n", "\n"), "\n") {
		parseGame(&game)
	}
	// parseGames(file)
}
