package main
type SnakeBody struct{
	X	int
	Y	int
	Xspeed	int
	Yspeed	int
	Body	int
};
func (sb *SnakeBody) ChangeDir(horizontal int, vertical int) {
	sb.Xspeed = horizontal;
	sb.Yspeed = vertical;
}
func (g *Game) Update() {
	g.snakeBody.X = (g.snakeBody.X + g.snakeBody.Xspeed) % g.width;
	if g.snakeBody.X < 0 {
		g.snakeBody.X = g.snakeBody.X + g.width;
	}
	g.snakeBody.Y = (g.snakeBody.Y + g.snakeBody.Yspeed) % g.height;
	if g.snakeBody.Y < 0 {
		g.snakeBody.Y = g.snakeBody.Y + g.height;
	}
} 
func (g *Game) GenSnake(){
	for i:=0 ; i< g.snakeBody.Body; i++ {
		posX:= g.snakeBody.X;
		posY:= g.snakeBody.Y;
		g.Screen.SetContent(posX, posY, ' ' , nil, g.style.snakeStyle);
		if g.snakeBody.Xspeed == 1 && g.snakeBody.Yspeed == 0{
			posX++;
		} else if g.snakeBody.Xspeed == -1 && g.snakeBody.Yspeed == 0{
			posX--;
		} else if g.snakeBody.Xspeed == 0 && g.snakeBody.Yspeed == 1{
			posY--;
		} else if g.snakeBody.Xspeed == 0 && g.snakeBody.Yspeed == -1{
			posY++;
		}
	}
	g.Update();
}
func (sb *SnakeBody) Grow(){
	sb.Body++;
}
func (g *Game) HasEaten() bool{
	if g.snakeBody.X == g.food.X && g.snakeBody.Y == g.food.Y{
		return true;
	}
	return false;
}
func (g *Game) Eat(){
	g.Screen.SetContent(g.food.X, g.food.Y, ' ', nil, g.style.defStyle )
}