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

	/* UP */
	/* 4 4 16 16
	   8 0 0 0
	   8 0 0 0
	   0 0 0 0 */
	var validBoardU [4][4]uint16

	validBoardU[0][0] = 4
	validBoardU[0][1] = 4
	validBoardU[0][2] = 16
	validBoardU[0][3] = 16

	validBoardU[1][0] = 8
	validBoardU[1][1] = 0
	validBoardU[1][2] = 0
	validBoardU[1][3] = 0

	validBoardU[2][0] = 0
	validBoardU[2][1] = 0
	validBoardU[2][2] = 0
	validBoardU[2][3] = 0

	validBoardU[3][0] = 0
	validBoardU[3][1] = 0
	validBoardU[3][2] = 0
	validBoardU[3][3] = 0

	newBoardU := simMove(board, "U")

	/* DOWN */
	/* 0 0 0 0
	   0 0 0 0
	   4 0 0 0
	   8 4 16 16 */
	var validBoardD [4][4]uint16

	validBoardD[0][0] = 0
	validBoardD[0][1] = 0
	validBoardD[0][2] = 0
	validBoardD[0][3] = 0

	validBoardD[1][0] = 0
	validBoardD[1][1] = 0
	validBoardD[1][2] = 0
	validBoardD[1][3] = 0

	validBoardD[2][0] = 4
	validBoardD[2][1] = 0
	validBoardD[2][2] = 0
	validBoardD[2][3] = 0

	validBoardD[3][0] = 8
	validBoardD[3][1] = 4
	validBoardD[3][2] = 16
	validBoardD[3][3] = 16

	newBoardD := simMove(board, "D")

	/* LEFT */
	/* 2 4 0 0
	   2 8 4 0
	   8 0 0 0
	   16 4 0 0 */
	var validBoardL [4][4]uint16

	validBoardL[0][0] = 2
	validBoardL[0][1] = 4
	validBoardL[0][2] = 0
	validBoardL[0][3] = 0

	validBoardL[1][0] = 2
	validBoardL[1][1] = 8
	validBoardL[1][2] = 4
	validBoardL[1][3] = 0

	validBoardL[2][0] = 8
	validBoardL[2][1] = 0
	validBoardL[2][2] = 0
	validBoardL[2][3] = 0

	validBoardL[3][0] = 16
	validBoardL[3][1] = 4
	validBoardL[3][2] = 0
	validBoardL[3][3] = 0

	newBoardL := simMove(board, "L")

	/* RIGHT */
	/* 0 0 2 4
	   0 2 8 4
	   0 0 0 8
	   0 0 16 4 */
	var validBoardR [4][4]uint16

	validBoardR[0][0] = 0
	validBoardR[0][1] = 0
	validBoardR[0][2] = 2
	validBoardR[0][3] = 4

	validBoardR[1][0] = 0
	validBoardR[1][1] = 2
	validBoardR[1][2] = 8
	validBoardR[1][3] = 4

	validBoardR[2][0] = 0
	validBoardR[2][1] = 0
	validBoardR[2][2] = 0
	validBoardR[2][3] = 8

	validBoardR[3][0] = 0
	validBoardR[3][1] = 0
	validBoardR[3][2] = 16
	validBoardR[3][3] = 4

	newBoardR := simMove(board, "R")

	if newBoardU != validBoardU {
		t.Errorf("UP Boards dont match\nStart:  %v\nResult: %v\nValid:  %v",
			board, newBoardU, validBoardU)
	}
	if newBoardD != validBoardD {
		t.Errorf("DOWN Boards dont match\nStart:  %v\nResult: %v\nValid:  %v",
			board, newBoardD, validBoardD)
	}
	if newBoardL != validBoardL {
		t.Errorf("LEFT Boards dont match\nStart:  %v\nResult: %v\nValid:  %v",
			board, newBoardL, validBoardL)
	}
	if newBoardR != validBoardR {
		t.Errorf("RIGHT Boards dont match\nStart:  %v\nResult: %v\nValid:  %v",
			board, newBoardR, validBoardR)
	}
}
