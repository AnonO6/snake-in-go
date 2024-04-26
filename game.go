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
}
func (g *Game) Run() {
	
	g.Screen.SetStyle(g.style.defStyle);
	for{
		g.Screen.Clear();
		g.GenSnake();
		g.Screen.SetContent(g.food.X, g.food.Y, '*', nil, g.style.foodStyle );
		if g.HasEaten() {
			g.Eat();
			g.snakeBody.Grow();
			g.food.GenFood(g.width, g.height);
		}
		time.Sleep(80 *time.Millisecond);
		g.Screen.Show();
	}
}