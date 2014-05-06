package main

import "testing"

func TestSimMove(t *testing.T) {
	/* 0 2 0 4
	   2 0 8 4
	   2 2 0 4
	   8 0 8 4 */
	var board [4][4]uint16
	board[0][0] = 0
	board[0][1] = 2
	board[0][2] = 0
	board[0][3] = 4

	board[1][0] = 2
	board[1][1] = 0
	board[1][2] = 8
	board[1][3] = 4

	board[2][0] = 2
	board[2][1] = 2
	board[2][2] = 0
	board[2][3] = 4

	board[3][0] = 8
	board[3][1] = 0
	board[3][2] = 8
	board[3][3] = 4

	/* 2 2 8 4
	   2 2 8 4
	   8 0 0 4
	   0 0 0 4 */
	var validBoard [4][4]uint16

	validBoard[0][0] = 4
	validBoard[0][1] = 4
	validBoard[0][2] = 16
	validBoard[0][3] = 16

	validBoard[1][0] = 8
	validBoard[1][1] = 0
	validBoard[1][2] = 0
	validBoard[1][3] = 0

	validBoard[2][0] = 0
	validBoard[2][1] = 0
	validBoard[2][2] = 0
	validBoard[2][3] = 0

	validBoard[3][0] = 0
	validBoard[3][1] = 0
	validBoard[3][2] = 0
	validBoard[3][3] = 0

	newBoard := simMove(board, "U")

	if newBoard != validBoard {
		t.Errorf("Boards dont match\nStart:  %v\nResult: %v\nValid:  %v",
			board, newBoard, validBoard)
	}
}
