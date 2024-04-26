package main

import "math/rand"
type Food struct{
	X	int
	Y	int
}
func (food *Food) GenFood(width int, height int){
	food.X = rand.Intn(width-1)
	food.Y = rand.Intn(height-1)
}
