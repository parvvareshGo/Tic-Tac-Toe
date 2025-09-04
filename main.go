package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printBoard(board [3][3]string) {
	fmt.Println("-------------")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("| %s ", board[i][j])
		}
		fmt.Println("|")
		fmt.Println("-------------")
	}
}

func checkWin(board [3][3]string, player string) bool {
	// Rows, columns
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	// Diagonals
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}

func checkDraw(board [3][3]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}

func playerMove(board *[3][3]string) {
	var row, col int
	for {
		fmt.Print("Enter row (0-2): ")
		fmt.Scan(&row)
		fmt.Print("Enter col (0-2): ")
		fmt.Scan(&col)
		if row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == " " {
			board[row][col] = "X"
			break
		} else {
			fmt.Println("Invalid move, try again.")
		}
	}
}

func computerMove(board *[3][3]string) {
	rand.Seed(time.Now().UnixNano())
	for {
		row := rand.Intn(3)
		col := rand.Intn(3)
		if board[row][col] == " " {
			board[row][col] = "O"
			fmt.Printf("Computer chose row %d, col %d\n", row, col)
			break
		}
	}
}

func main() {
	var board [3][3]string
	// Initialize board with empty spaces
	for i := range board {
		for j := range board[i] {
			board[i][j] = " "
		}
	}

	fmt.Println("Tic-Tac-Toe Game: You (X) vs Computer (O)")
	for {
		printBoard(board)
		playerMove(&board)
		if checkWin(board, "X") {
			printBoard(board)
			fmt.Println("You win!")
			break
		}
		if checkDraw(board) {
			printBoard(board)
			fmt.Println("Draw!")
			break
		}

		computerMove(&board)
		if checkWin(board, "O") {
			printBoard(board)
			fmt.Println("Computer wins!")
			break
		}
		if checkDraw(board) {
			printBoard(board)
			fmt.Println("Draw!")
			break
		}
	}
}
