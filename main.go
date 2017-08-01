package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Let's play Tic-Tac-Toe")
	fmt.Println("-------------------------------")

	board, player := newGame()

	for {

		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println("User input was invalid")
			continue
		}

		switch char {
		case '1':
			selectTile(&board, &player, 2, 0)
			break
		case '2':
			selectTile(&board, &player, 2, 1)
			break
		case '3':
			selectTile(&board, &player, 2, 2)
			break
		case '4':
			selectTile(&board, &player, 1, 0)
			break
		case '5':
			selectTile(&board, &player, 1, 1)
			break
		case '6':
			selectTile(&board, &player, 1, 2)
			break
		case '7':
			selectTile(&board, &player, 0, 0)
			break
		case '8':
			selectTile(&board, &player, 0, 1)
			break
		case '9':
			selectTile(&board, &player, 0, 2)
			break
		}
	}

}

func selectTile(board *[3][]string, player *int, row int, col int) {
	p := *player
	fmt.Printf("Player %d selected a tile!\n", p)

	switch p {
	case 1:
		board[row][col] = "X"
		break
	case 0:
		board[row][col] = "O"
		break
	}

	if gameEnded(*board) {
		fmt.Printf("Player %d won! \n", p)
		for i := 0; i < len(board); i++ {
			fmt.Printf("%s\n", strings.Join(board[i], " "))
		}
		*board, *player = newGame()
		return
	}

	*player = *player ^ 1

	fmt.Println("The current board is:")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println("-------------------------------")
	fmt.Printf("Player %d is up.\n", p)
}

func newGame() (board [3][]string, player int) {
	fmt.Println("-------------------------------")
	fmt.Println("New Game!")
	fmt.Println("Player 0 is up")
	return [3][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}, 0
}

func gameEnded(board [3][]string) bool {
	return diagonalWin(board) || horizontalWin(board) || verticalWin(board)
}

func horizontalWin(board [3][]string) bool {
	for i := 0; i < len(board); i++ {
		if board[i][0] == "_" {
			return false
		}
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return true
		}
	}
	return false
}

func verticalWin(board [3][]string) bool {
	for i := 0; i < len(board); i++ {
		if board[0][i] == "_" {
			return false
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return true
		}
	}
	return false
}

func diagonalWin(board [3][]string) bool {
	for i := 0; i < len(board); i++ {
		if board[0][i%3] == "_" {
			return false
		}
		if board[0][i%3] == board[1][(i+1)%3] && board[1][(i+1)%3] == board[2][(i+2)%3] {
			return true
		}
	}
	return false
}
