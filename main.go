package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/BurntSushi/xgb/xproto"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xgraphics"
)

var offWhite = xgraphics.BGRA{239, 248, 250, 255} //color of page background
var bgGrey = xgraphics.BGRA{160, 173, 187, 255}   //color of board background
var tile0 = xgraphics.BGRA{179, 192, 204, 255}    //color of empty tile
var tile2 = xgraphics.BGRA{218, 228, 238, 255}    //color of 2 tile
var tile4 = xgraphics.BGRA{200, 224, 237, 255}    //4 tile
var tile8 = xgraphics.BGRA{121, 177, 242, 255}    //8 tile
var tile16 = xgraphics.BGRA{99, 149, 245, 255}    //16 tile
var tile32 = xgraphics.BGRA{95, 124, 246, 255}    //32 tile
var tile64 = xgraphics.BGRA{59, 94, 246, 255}     //64 tile
var tile128 = xgraphics.BGRA{114, 207, 237, 255}  //128 tile
var tile256 = xgraphics.BGRA{97, 204, 237, 255}   //256 tile
var tile512 = xgraphics.BGRA{80, 200, 237, 255}   //512 tile
var tile1024 = xgraphics.BGRA{63, 197, 237, 255}  //1024 tile
var tile2048 = xgraphics.BGRA{46, 194, 237, 255}  //2048 tile
var tile4096 = xgraphics.BGRA{50, 58, 60, 255}    //4096+ tile, i think 8192 and onward is also the same color

func colorNum(c color.Color) (uint16, error) {
	switch c {
	case tile0:
		return 0, nil
	case tile2:
		return 2, nil
	case tile4:
		return 4, nil
	case tile8:
		return 8, nil
	case tile16:
		return 16, nil
	case tile32:
		return 32, nil
	case tile64:
		return 64, nil
	case tile128:
		return 128, nil
	case tile256:
		return 256, nil
	case tile512:
		return 512, nil
	case tile1024:
		return 1024, nil
	case tile2048:
		return 2048, nil
	case tile4096:
		return 4096, nil
	}
	return 0, fmt.Errorf("Unrecognized color: %v", c)
}

func exploreMoves(board [4][4]uint16, moves []byte, depth int) ([]byte, int) {
	if depth == 0 {
		//return calcMatches(board)
		score := 0
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if board[i][j] != 0 {
					score += int(board[i][j]) - 7
				}
			}
		}
		return moves, score
	}
	var umoves, dmoves, lmoves, rmoves []byte
	up, down, left, right := 0, 0, 0, 0
	uboard := simMove(board, "U")
	dboard := simMove(board, "D")
	lboard := simMove(board, "L")
	rboard := simMove(board, "R")
	if uboard == board {
		up = -100
	} else {
		umoves, up = exploreMoves(uboard, moves, depth-1)
	}
	if dboard == board {
		down = -100
	} else {
		dmoves, down = exploreMoves(dboard, moves, depth-1)
	}
	if lboard == board {
		left = -100
	} else {
		lmoves, left = exploreMoves(lboard, moves, depth-1)
	}
	if rboard == board {
		right = -100
	} else {
		rmoves, right = exploreMoves(rboard, moves, depth-1)
	}
	fmt.Println("Moves:", up, down, left, right, " -- ", moves, depth)
	if up >= down {
		if up >= left {
			if up >= right {
				return append(umoves, 'U'), up
			}
			return append(rmoves, 'R'), right
		}
		if left >= right {
			return append(lmoves, 'L'), left
		}
		return append(rmoves, 'R'), right
	}
	if down >= left {
		if down >= right {
			return append(dmoves, 'D'), down
		}
		return append(rmoves, 'R'), right
	}
	if left >= right {
		return append(lmoves, 'L'), left
	}
	return append(rmoves, 'R'), right
}

func findMove(board [4][4]uint16) []byte {
	depth := 2
	moves := make([]byte, depth)
	moves, _ = exploreMoves(board, moves, depth)
	//fmt.Println(calcMatches(board))
	for i, j := 0, len(moves)-1; i < j; i, j = i+1, j-1 {
		moves[i], moves[j] = moves[j], moves[i]
	}
	return moves
}

func simMove(board [4][4]uint16, move string) [4][4]uint16 {
	switch move {
	case "U":
		for i := 0; i < 4; i++ {
			//shift tiles
			free := 4
			for j := 0; j < 4; j++ { //find first empty tile
				if board[j][i] == 0 {
					free = j
					break
				}
			}
			for j := free + 1; j < 4; j++ { //move each non-empty tile to first empty tile
				if board[j][i] == 0 {
					continue
				}
				board[free][i] = board[j][i]
				board[j][i] = 0
				free++
			}
			//merge tiles
			for j := 0; j < 3; j++ {
				if board[j][i] == board[j+1][i] { //marge matching tiles
					board[j][i] *= 2
					for k := j + 1; k < 3; k++ { //shift following tiles
						board[k][i] = board[k+1][i]
					}
					board[3][i] = 0
				}
			}
		}
	case "D":
		for i := 0; i < 4; i++ {
			//shift tiles
			free := -1
			for j := 3; j >= 0; j-- { //find first empty tile
				if board[j][i] == 0 {
					free = j
					break
				}
			}
			for j := free - 1; j >= 0; j-- { //move each non-empty tile to first empty tile
				if board[j][i] == 0 {
					continue
				}
				board[free][i] = board[j][i]
				board[j][i] = 0
				free--
			}
			//merge tiles
			for j := 3; j >= 1; j-- {
				if board[j][i] == board[j-1][i] { //marge matching tiles
					board[j][i] *= 2
					for k := j - 1; k >= 1; k-- { //shift following tiles
						board[k][i] = board[k-1][i]
					}
					board[0][i] = 0
				}
			}
		}
	case "L":
		for i := 0; i < 4; i++ {
			//shift tiles
			free := 4
			for j := 0; j < 4; j++ { //find first empty tile
				if board[i][j] == 0 {
					free = j
					break
				}
			}
			for j := free + 1; j < 4; j++ { //move each non-empty tile to first empty tile
				if board[i][j] == 0 {
					continue
				}
				board[i][free] = board[i][j]
				board[i][j] = 0
				free++
			}
			//merge tiles
			for j := 0; j < 3; j++ {
				if board[i][j] == board[i][j+1] { //marge matching tiles
					board[i][j] *= 2
					for k := j + 1; k < 3; k++ { //shift following tiles
						board[i][k] = board[i][k+1]
					}
					board[i][3] = 0
				}
			}
		}
	case "R":
		for i := 0; i < 4; i++ {
			//shift tiles
			free := -1
			for j := 3; j >= 0; j-- { //find first empty tile
				if board[i][j] == 0 {
					free = j
					break
				}
			}
			for j := free - 1; j >= 0; j-- { //move each non-empty tile to first empty tile
				if board[i][j] == 0 {
					continue
				}
				board[i][free] = board[i][j]
				board[i][j] = 0
				free--
			}
			//merge tiles
			for j := 3; j >= 1; j-- {
				if board[i][j] == board[i][j-1] { //marge matching tiles
					board[i][j] *= 2
					for k := j - 1; k >= 1; k-- { //shift following tiles
						board[i][k] = board[i][k-1]
					}
					board[i][0] = 0
				}
			}
		}
	}
	return board
}

func main() {
	x11, err := xgbutil.NewConn() //connect to X
	if err != nil {
		log.Fatal(err)
	}

	var board [4][4]uint16

	img, err := xgraphics.NewDrawable(x11, xproto.Drawable(x11.RootWin())) //get screen
	if err != nil {
		log.Fatal(err)
	}
	imgRect := img.Bounds()

	boardX, boardY := 0, 0

OuterLoop: //search screen for upper left corner of board
	for x := imgRect.Min.X; x < imgRect.Max.X; x++ {
		for y := imgRect.Min.Y; y < imgRect.Max.Y; y++ {
			if img.At(x, y) == bgGrey { //board might be found
				x0, x1, x2, x3, x4, x5 := x, x+1, x+2, x+3, x+4, x+5
				y0, y1, y2, y3, y4, y5 := y-5, y-4, y-3, y-2, y-1, y
				if img.At(x0, y0) == offWhite && //check corner pixel and board pixel colors
					img.At(x0, y1) == offWhite &&
					img.At(x0, y2) == offWhite &&
					img.At(x1, y0) == offWhite &&
					img.At(x2, y0) == offWhite &&
					img.At(x5, y0) == bgGrey &&
					img.At(x5, y1) == bgGrey &&
					img.At(x5, y2) == bgGrey &&
					img.At(x5, y3) == bgGrey &&
					img.At(x5, y4) == bgGrey &&
					img.At(x5, y5) == bgGrey &&
					img.At(x0, y5) == bgGrey &&
					img.At(x1, y5) == bgGrey &&
					img.At(x2, y5) == bgGrey &&
					img.At(x3, y5) == bgGrey &&
					img.At(x4, y5) == bgGrey &&
					img.At(x5, y5) == bgGrey {
					boardX, boardY = x0, y0
					fmt.Println("Board corner at", boardX, ",", boardY)
					break OuterLoop
				}
			}
		}
	}

	tileX, tileY := boardX+63, boardY+39
	lastBoard := board
ImgLoop:
	for i := 0; true; i++ {
		img, err := xgraphics.NewDrawable(x11, xproto.Drawable(x11.RootWin())) //get screen
		if err != nil {
			log.Fatal(err)
		}
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				board[y][x], err = colorNum(img.At(tileX+(121*x), tileY+(121*y)))
				if err != nil {
					fmt.Println(err, " at ", tileX+(121*x), tileY+(121*y))
					time.Sleep(1 * time.Second)
					continue ImgLoop
				}
				//fmt.Print(board[x][y], " ")
			}
			//fmt.Println()
		}
		if lastBoard != board {
			fmt.Println(string(findMove(board)), i)
		}
		lastBoard = board
		time.Sleep(1 * time.Second)
	}
}
