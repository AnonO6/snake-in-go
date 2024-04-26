package main

import (
	"time"

	"github.com/gdamore/tcell"
)
type Game struct{
	Screen	tcell.Screen
	snakeBody	SnakeBody
	food	Food
	width	int
	height	int
	style	Style
};
type Style struct{
	defStyle	tcell.Style
	snakeStyle	tcell.Style
	foodStyle	tcell.Style
	scoreStyle	tcell.Style
}
func (g *Game) Run() {
	
	g.Screen.SetStyle(g.style.defStyle);
	for{
		g.Screen.Clear();
		g.Score();
		g.GenSnake();
		for i:=0 ; i<3 ; i++{
			g.Screen.SetContent(g.food.X[i], g.food.Y[i], '\u25CF', nil, g.style.foodStyle );
		}
		if index:= g.HasEaten(); index != -1 {
			g.Eat(index);
			g.snakeBody.Grow();
			g.GenFood(g.width, g.height, index);
		}
		time.Sleep(160 *time.Millisecond);
		if g.checkCollision(g.snakeBody.Body[g.snakeBody.length-1]){
			break;
		}
		g.Screen.Show();
	}
	g.GameOver();
	g.Screen.Show();
}