package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

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

func colorNum(c color.Color) uint16 {
	switch c {
	case tile0:
		return 0
	case tile2:
		return 2
	case tile4:
		return 4
	case tile8:
		return 8
	case tile16:
		return 16
	case tile32:
		return 32
	case tile64:
		return 64
	case tile128:
		return 128
	case tile256:
		return 256
	case tile512:
		return 512
	case tile1024:
		return 1024
	case tile2048:
		return 2048
	case tile4096:
		return 4096
	}
	log.Panic("Unknown color!: ", c)
	return 0
}

func exploreMoves(board [4][4]uint16, depth int) (string, int) {
	if depth == 0 {
		//return calcMatches(board)
		score := 0
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if board[i][j] != 0{
					score += int((math.Log2(float64(board[i][j])) - 2))
				}
			}
		}
		return "", score
	}
	_, up := exploreMoves(simMove(board, "U"), depth-1)
	_, down := exploreMoves(simMove(board, "D"), depth-1)
	_, left := exploreMoves(simMove(board, "L"), depth-1)
	_, right := exploreMoves(simMove(board, "R"), depth-1)
	fmt.Println("Moves:", up, down, left, right)
	if up >= down {
		if up >= left {
			if up >= right {
				return "U", up
			}
			return "R", right
		}
		if left >= right {
			return "L", left
		}
		return "R", right
	}
	if down >= left {
		if down >= right {
			return "D", down
		}
		return "R", right
	}
	if left >= right {
		return "L", left
	}
	return "R", right
}

func calcMatches(board [4][4]uint16) (string, int) {
	rowMatches, colMatches := 0, 0

	//find matches made by a move row-wise (left or right)
	for i := 0; i < 4; i++ {
		lastSeen := board[i][0]
		for j := 1; j < 4; j++ {
			if board[i][j] == 0 {
				continue
			}
			if lastSeen == 0 {
				lastSeen = board[i][j]
				continue
			}
			if lastSeen == board[i][j] {
				rowMatches++
				lastSeen = 0
			} else {
				lastSeen = board[i][j]
			}
		}
	}

	//find matches made by a move column-wise (up or down)
	for i := 0; i < 4; i++ {
		lastSeen := board[0][i]
		for j := 1; j < 4; j++ {
			if board[j][i] == 0 {
				continue
			}
			if lastSeen == 0 {
				lastSeen = board[j][i]
				continue
			}
			if lastSeen == board[j][i] {
				colMatches++
				lastSeen = 0
			} else {
				lastSeen = board[j][i]
			}
		}
	}

	//fmt.Println(rowMatches)
	//fmt.Println(colMatches)
	if rowMatches > colMatches {
		return "L", rowMatches
	}
	return "U", colMatches
}

func findMove(board [4][4]uint16) string {
	move, _ := exploreMoves(board, 3)
	//fmt.Println(calcMatches(board))
	return move
}

func simMove(board [4][4]uint16, move string) [4][4]uint16 {
	switch move{
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
			for j := free+1; j < 4; j++ { //move each non-empty tile to first empty tile
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
					for k := j+1; k < 3; k++{ //shift following tiles
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
			for j := free-1; j >= 0; j-- { //move each non-empty tile to first empty tile
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
					for k := j-1; k >= 1; k--{ //shift following tiles
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
			for j := free+1; j < 4; j++ { //move each non-empty tile to first empty tile
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
					for k := j+1; k < 3; k++{ //shift following tiles
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
			for j := free-1; j >= 0; j-- { //move each non-empty tile to first empty tile
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
					for k := j-1; k >= 1; k--{ //shift following tiles
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

	tileX, tileY := boardX+20, boardY+20
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			board[x][y] = colorNum(img.At(tileX+(121*y), tileY+(121*x)))
			fmt.Print(board[x][y], " ")
		}
		fmt.Println()
	}
	fmt.Println(findMove(board))
}
