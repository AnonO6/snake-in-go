package main

import "math/rand"
type Food struct{
	X	[3]int
	Y	[3]int
}
func (g *Game) GenFood(width int, height int, index int){
		g.food.X[index] = rand.Intn(width-1)
		g.food.Y[index] = rand.Intn(height-1)
}
