package main

import (
	"fmt"
	"log"
	"image/color"

	"github.com/BurntSushi/xgb/xproto"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xgraphics"
)

var offWhite = xgraphics.BGRA{239, 248, 250, 255} //color of page background
var bgGrey = xgraphics.BGRA{160, 173, 187, 255}   //color of board background
var tile0 = xgraphics.BGRA{179, 192, 204, 255} //color of empty tile
var tile2 = xgraphics.BGRA{218, 228, 238, 255} //color of 2 tile
var tile4 = xgraphics.BGRA{200, 224, 237, 255} //4 tile
var tile8 = xgraphics.BGRA{121, 177, 242, 255} //8 tile
var tile16 = xgraphics.BGRA{99, 149, 245, 255} //16 tile
var tile32 = xgraphics.BGRA{95, 124, 246, 255} //32 tile
var tile64 = xgraphics.BGRA{59, 94, 246, 255} //64 tile
var tile128 = xgraphics.BGRA{114, 207, 237, 255} //128 tile
var tile256 = xgraphics.BGRA{97, 204, 237, 255} //256 tile
var tile512 = xgraphics.BGRA{80, 200, 237, 255} //512 tile
var tile1024 = xgraphics.BGRA{63, 197, 237, 255} //1024 tile
var tile2048 = xgraphics.BGRA{46, 194, 237, 255} //2048 tile
var tile4096 = xgraphics.BGRA{50, 58, 60, 255} //4096+ tile, i think 8192 and onward is also the same color

func colorNum(c color.Color) uint16 {
	switch c {
	case tile0: return 0
	case tile2: return 2
	case tile4: return 4
	case tile8: return 8
	case tile16: return 16
	case tile32: return 32
	case tile64: return 64
	case tile128: return 128
	case tile256: return 256
	case tile512: return 512
	case tile1024: return 1024
	case tile2048: return 2048
	case tile4096: return 4096
	}
	log.Panic("Unknown color!: ", c)
	return 0
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
			fmt.Print(board[x][y])
		}
		fmt.Println()
	}
}
