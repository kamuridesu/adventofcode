package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	color    string
	limit    int
	quantity int
}

type Set struct {
	red   Cube
	blue  Cube
	green Cube
}

func (s *Set) addCube(cube *Cube) {
	switch cube.color {
	case "red":
		s.red = *cube
	case "blue":
		s.blue = *cube
	case "green":
		s.green = *cube
	}

}

type Game struct {
	id   int
	sets []Set
}

func NewCube(color string, quantity int) *Cube {
	var limit int
	switch color {
	case "blue":
		limit = 14
	case "red":
		limit = 12
	case "green":
		limit = 13
	default:
		limit = 0
	}
	return &Cube{
		color:    color,
		limit:    limit,
		quantity: quantity,
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

func parseGame(input *string) *Game {
	gameData := strings.Split(*input, ":")
	gameId, err := strconv.Atoi(strings.Split(gameData[0], " ")[1])
	if err != nil {
		log.Fatal(err)
	}
	game := Game{id: gameId}
	sets := strings.Split(gameData[1], ";")
	for _, set := range sets {
		set = strings.TrimSpace(set)
		items := strings.Split(set, ",")
		setStruct := Set{}
		for _, item := range items {
			item := strings.TrimSpace(item)
			quantityAndColor := strings.Split(item, " ")
			quantityNumber, err := strconv.Atoi(quantityAndColor[0])
			if err != nil {
				log.Fatal(err)
			}
			cube := NewCube(quantityAndColor[1], quantityNumber)
			setStruct.addCube(cube)
		}
		game.sets = append(game.sets, setStruct)
	}
	return &game
}

func (g *Game) checkIfCubeQuantityIsBiggerThanLimit() bool {
	isValid := true
	for _, set := range g.sets {
		if (set.blue.color != "" && set.blue.quantity > set.blue.limit) ||
			(set.green.color != "" && set.green.quantity > set.green.limit) ||
			(set.red.color != "" && set.red.quantity > set.red.limit) {
			isValid = false
		}
	}

	return isValid
}

func (g *Game) findMinimunNecessaryCubes() int {
	minimumRed := 0
	minimumBlue := 0
	minimumGren := 0
	for _, set := range g.sets {
		if set.blue.quantity > minimumBlue {
			minimumBlue = set.blue.quantity
		}
		if set.green.quantity > minimumGren {
			minimumGren = set.green.quantity
		}
		if set.red.quantity > minimumRed {
			minimumRed = set.red.quantity
		}
	}
	return minimumBlue * minimumGren * minimumRed
}

func sumIds(g []*Game) int {
	total := 0
	for _, game := range g {
		total += game.id
	}
	return total
}

func main() {
	file := readFileAsString()
	var games []*Game
	totalMinimumPowerOfAllGames := 0
	for _, game := range strings.Split(strings.ReplaceAll(*file, "\r\n", "\n"), "\n") {
		g := parseGame(&game)
		isValid := g.checkIfCubeQuantityIsBiggerThanLimit()
		if isValid {
			games = append(games, g)
		}
		gameMinimumPower := g.findMinimunNecessaryCubes()
		totalMinimumPowerOfAllGames += gameMinimumPower
	}
	sumOfIds := sumIds(games)
	fmt.Printf("Total sum of valid Game Ids: %v\n", sumOfIds)                         // Solution 1
	fmt.Printf("Total minimum power of all games: %v\n", totalMinimumPowerOfAllGames) // Solution 2
}
