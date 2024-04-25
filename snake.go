package main
type SnakeBody struct{
	X	int
	Y	int
	Xspeed	int
	Yspeed	int
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