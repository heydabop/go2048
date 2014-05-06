package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/xgb/xproto"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xgraphics"
)

var offWhite = xgraphics.BGRA{239, 248, 250, 255} //color of page background
var bgGrey = xgraphics.BGRA{160, 173, 187, 255}   //color of board background

func findMove(board [4][4]uint16) byte {
	rowMatches, colMatches := 0, 0

	//find matches made by a move row-wise (left or right)
	for i := 0; i < 4; i++ {
		lastSeen := board[i][0]
		for j := 0; j < 4; j++ {
			if lastSeen == 0 {
				lastSeen = board[i][j]
				continue
			}
			if board[i][j] == 0 {
				continue
			}
			if lastSeen == board[i][j] {
				rowMatches++
			}
			lastSeen = board[i][j]
		}
	}

	//find matches made by a move column-wise (up or down)
	for i := 0; i < 4; i++ {
		lastSeen := board[0][i]
		for j := 0; j < 4; j++ {
			if lastSeen == 0 {
				lastSeen = board[j][i]
				continue
			}
			if board[j][i] == 0 {
				continue
			}
			if lastSeen == board[j][i] {
				rowMatches++
			}
			lastSeen = board[j][i]
		}
	}

	if rowMatches > colMatches {
		return 'U'
	}
	return 'L'
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
					fmt.Println("Board corner at ", x0, ",", y0)
					break OuterLoop
				}
			}
		}
	}
}
