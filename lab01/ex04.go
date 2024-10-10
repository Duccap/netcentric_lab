package main

import (
	"fmt"
	"math/rand"
	"strconv"
)


func initializeBoard(board [][]string, rows int, cols int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			board[i][j] = "."
		}
	}
}

func placeMines(board [][]string, rows int, cols int, mines int) {
	placedMines := 0
	for placedMines < mines {
		mineRow := rand.Intn(rows)
		mineCol := rand.Intn(cols)
		if board[mineRow][mineCol] != "*" {
			board[mineRow][mineCol] = "*"
			placedMines++
		}
	}
}

func markBoard(board [][]string, rows int, cols int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == "." {
				adjacentMines := countAdjacentMines(board, i, j, rows, cols)
				if adjacentMines > 0 {
					board[i][j] = strconv.Itoa(adjacentMines)
				} else {
					board[i][j] = "."
				}
			}
		}
	}
}

func countAdjacentMines(board [][]string, row int, col int, rows int, cols int) int {
	count := 0
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if r >= 0 && r < rows && c >= 0 && c < cols && board[r][c] == "*" {
				count++
			}
		}
	}
	return count
}

func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j])
		}
		fmt.Println()
	}
}

func generateMinesweeperBoard(rows int, cols int, mines int) [][]string {
	board := make([][]string, rows)
	for i := range board {
		board[i] = make([]string, cols)
	}
	initializeBoard(board, rows, cols)
	placeMines(board, rows, cols, mines)
	markBoard(board, rows, cols)
	return board
}

func main() {
	board := generateMinesweeperBoard(20, 25, 99) 
	printBoard(board) 
}
