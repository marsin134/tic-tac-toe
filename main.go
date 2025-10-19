package main

import (
	"errors"
	"fmt"
)

var (
	CoordsNotCorrectError = errors.New("Coords not correct")
	CoordsOccupiedCell    = errors.New("Coordinates of the occupied cell")
)

func draw(lines [3][3]string) {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			fmt.Printf("| %s ", lines[i][j])
		}
		fmt.Printf("|\n")
	}
}

func correctInputCoords(coords [2]int, field [3][3]string) ([2]int, error) {
	if coords[0] < 1 || coords[1] < 1 || coords[0] > 3 || coords[1] > 3 {
		return [2]int{0, 0}, CoordsNotCorrectError
	}
	if field[coords[0]-1][coords[1]-1] != "." {
		return [2]int{0, 0}, CoordsOccupiedCell
	}

	return coords, nil
}

func inputCoords(field [3][3]string) ([2]int, error) {
	coords := [2]int{0, 0}

	fmt.Scan(&coords[0], &coords[1])

	return correctInputCoords(coords, field)
}

func game() {
	field := [3][3]string{{".", "X", "."}, {"O", ".", "."}, {".", ".", "."}}
	indexPlayer := 0

	var commandUserInput string

	fmt.Print("Start game. The first crosses go")
	fmt.Println("Сommands: 'c number_column number_line' example: 'c 1 2'")
	fmt.Println("\t's' - stop game")

	for {
		switch commandUserInput {
		case "c":
			coords, err := inputCoords(field)
			if err != nil {
				fmt.Println(err)
			} else {
				if indexPlayer == 0 {
					field[coords[0]-1][coords[1]-1] = "X"
				} else {
					field[coords[0]-1][coords[1]-1] = "O"
				}

				indexPlayer = (indexPlayer + 1) % 2

				draw(field)
			}

		case "s":
			return

		default:
			fmt.Println("Unknown command")
		}

		fmt.Printf("Input command: ")
		fmt.Scan(&commandUserInput)
	}
}

func main() {
	game()
}
