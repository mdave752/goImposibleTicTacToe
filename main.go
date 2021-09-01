package main

import (
	"fmt"
	"strings"
)

func placeIfTwoInRow(win *[8][3]int, board *[9]string, val string) bool {
	for _, v := range win {
		t := board[v[0]] + board[v[1]] + board[v[2]]
		if strings.Count(t, val) > 1 {
			m := strings.Index(t, " ")
			if m >= 0 {
				board[v[m]] = "O"
				return true
			}
		}
	}
	return false
}

func computerPlays(win *[8][3]int, board *[9]string, marker *[2]string) {
	// checks if value in middle, if not, AI puts piece in middle.
	if board[4] == " " {
		board[4] = marker[1]
		return
	}
	if placeIfTwoInRow(win, board, marker[1]) {
		return
	}
	if placeIfTwoInRow(win, board, marker[0]) {
		return
	}
	// had to make this, becuase AI lost if you put pieces on sides, and and then got two win conditions
	checkForSides := [2][3]int{
		{5, 7, 8},
		{1, 5, 2},
	}
	for _, v := range checkForSides {
		if board[v[0]] == "X" && board[v[1]] == "X" && board[v[2]] != "O" {
			board[v[2]] = marker[1]
			return
		}
	}

	for _, v := range win {
		t := board[v[0]] + board[v[1]] + board[v[2]]
		if (board[v[0]] == " " || board[v[0]] == marker[1]) &&
			(board[v[1]] == " " || board[v[1]] == marker[1]) &&
			(board[v[2]] == " " || board[v[2]] == marker[1]) {
			m := strings.Index(t, " ")
			board[v[m]] = marker[1]
			return
		}
	}
	// dublicate of func Check free space
	for i, x := range board {
		if x == " " {
			board[i] = marker[1]
			return
		}
	}

}

func checkWinCondition(win *[8][3]int, board *[9]string, marker *[2]string) bool {
	for _, v := range win {
		for i := 0; i < len(marker); i++ {
			if board[v[0]] == marker[i] &&
				board[v[1]] == marker[i] && board[v[2]] == marker[i] {
				return true
			}
		}
	}
	return false
}

func checkFreeSpace(board *[9]string) bool {
	for _, v := range board {
		if v == " " {
			return true
		}
	}
	return false
}

func printBoard(board *[9]string) {
	fmt.Println()
	for i := 0; i < len(board); i += 3 {
		fmt.Println("|", board[i], "|", board[i+1], "|", board[i+2], "|")
	}
}

func playAgain() bool {
	var playAgain string
	fmt.Println("\nDo you want to play again y/n ?")
	fmt.Scanln(&playAgain)
	ifYes := strings.ToLower(playAgain)
	return strings.HasPrefix(ifYes, "y")

}

func placeMarkerX(input int, board *[9]string) bool {
	if input < 1 || input > 9 {
		fmt.Println("Please pick a number between 1-9")
		return false
	}
	if board[input-1] != " " {
		return false
	}
	return true
}
func main() {
	win := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}
	var createBoard [9]string
	for i := 0; i < len(createBoard); i++ {
		createBoard[i] = " "
	}

	board := createBoard

	var marker [2]string
	marker[0] = "X"
	marker[1] = "O"

	var input int

	var winsX bool
	var winsAI bool

	for {
		fmt.Print("\nPlease pick a place to put ", marker[0], " from 1-9 : ")
		fmt.Scanln(&input)

		if placeMarkerX(input, &board) {
			board[input-1] = marker[0]
		} else {
			continue
		}

		winsX = checkWinCondition(&win, &board, &marker)
		// computer calculates and places the marker O
		if !winsX {
			computerPlays(&win, &board, &marker)
			winsAI = checkWinCondition(&win, &board, &marker)
		}

		noSpace := checkFreeSpace(&board)

		printBoard(&board)

		if winsX || winsAI || !noSpace {
			if winsX {
				fmt.Println("\nX Won the game!")
			} else if winsAI {
				fmt.Println("\nO Won the game!")
			} else {
				fmt.Println("\nDraw!")
			}
			if playAgain() {
				board = createBoard
				continue
			} else {
				break
			}

		}
	}
}
