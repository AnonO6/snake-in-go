package main

import (
	"strconv"

	"github.com/gdamore/tcell"
)
func (g *Game) Score (){
	text:= "SCORE "+ strconv.Itoa(g.snakeBody.length-1);

	drawText(g.Screen, 1, 1, g.width, g.height, text);
}
func (g *Game) GameOver (){
	drawText(g.Screen, g.width/2, g.height/2,g.width, g.height, "GAME OVER");
}
func drawText(s tcell.Screen, x1, y1, x2, y2 int, text string) {
	row := y1
	col := x1
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	for _, r := range text {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}