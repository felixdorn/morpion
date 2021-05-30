package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	PlayerOneWon = iota
	PlayerTwoWon
	Playing
)

func main() {
	fmt.Println("Welcome to my yet to be tic-tac-toe game!")

	board := [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	var gameState = Playing
	var firstPlayerPlaying = true

	reader := bufio.NewReader(os.Stdin)

	for gameState == Playing {
		if firstPlayerPlaying {
			fmt.Println("It's player one's turn!")
		} else {
			fmt.Println("It's player two's turn!")
		}

		fmt.Print("Enter a board position : ")

		input, _ := reader.ReadString('\n')

		rawPosX, _ := strconv.ParseUint(input[0:1], 10, 4)
		posY := clamp(rawPosX-1, 0, 3)
		rawPosY, _ := strconv.ParseUint(input[1:2], 10, 4)
		posX := clamp(rawPosY-1, 0, 3)

		if board[posY][posX] != 0 {
			fmt.Println("Case already filled.")
			continue
		}

		if firstPlayerPlaying {
			board[posY][posX] = 1
		} else {
			board[posY][posX] = 10
		}

		printBoard(board)

		gameState = getGameState(board)
		firstPlayerPlaying = !firstPlayerPlaying
	}

	if gameState == PlayerOneWon {
		fmt.Println("Player one won!")
	}

	if gameState == PlayerTwoWon {
		fmt.Println("Player two won!")
	}

}

func printBoard(board [3][3]int) {
	fmt.Println("  1 2 3")
	for i, rows := range board {
		prettyColumn := ""
		for _, column := range rows {
			if column == 1 {
				prettyColumn += "x "
			} else if column == 10 {
				prettyColumn += "o "
			} else {
				prettyColumn += "_ "
			}
		}

		fmt.Println(i+1, prettyColumn)
	}

}

func getGameState(board [3][3]int) int {
	for i := 0; i < len(board); i++ {
		column := board[i]

		if sum(column[:]) == 30 {
			return PlayerTwoWon
		}

		if sum(column[:]) == 3 {
			return PlayerOneWon
		}
	}

	sumOfRows := [3]int{
		board[0][0] + board[1][0] + board[2][0],
		board[0][1] + board[1][1] + board[2][1],
		board[0][2] + board[1][2] + board[2][2],
	}

	if contains(sumOfRows[:], 3) {
		return PlayerOneWon
	}

	if contains(sumOfRows[:], 30) {
		return PlayerTwoWon
	}

	sumOfDiagonals := [2]int{
		board[0][0] + board[1][1] + board[2][2],
		board[2][0] + board[1][1] + board[0][2],
	}

	if contains(sumOfDiagonals[:], 3) {
		return PlayerOneWon
	}

	if contains(sumOfDiagonals[:], 30) {
		return PlayerTwoWon
	}

	return Playing
}

func clamp(x uint64, min uint64, max uint64) uint64 {
	if x < min {
		return min
	}

	if x > max {
		return max
	}

	return x
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func contains(arr []int, comp int) bool {
	for _, a := range arr {
		if a == comp {
			return true
		}
	}
	return false
}
