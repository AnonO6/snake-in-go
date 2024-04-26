package main
type SnakeBody struct{
	Xspeed	int
	Yspeed	int
	Body	[][2]int
	length	int
};
func (sb *SnakeBody) Enqueue(part [2]int){
	sb.Body = append(sb.Body, part)
}
func (sb *SnakeBody) Dequeue(){
	sb.Body = sb.Body[1:]
}
func (sb *SnakeBody) ChangeDir(horizontal int, vertical int) {
	sb.Xspeed = horizontal;
	sb.Yspeed = vertical;
}
func (g *Game) Update() {
	head:= g.snakeBody.Body[g.snakeBody.length-1]
	X:= head[0]
	Y:= head[1]

	X = (X + g.snakeBody.Xspeed) % g.width;
	if X < 0 {
		X = X + g.width;
	}
	Y = (Y + g.snakeBody.Yspeed) % g.height;
	if Y < 0 {
		Y = Y + g.height;
	}
	g.snakeBody.Enqueue([2]int {X,Y});
	g.snakeBody.Dequeue();
} 
func (g *Game) GenSnake(){
	for i:=0 ; i< g.snakeBody.length; i++ {
		g.Screen.SetContent(g.snakeBody.Body[i][0], g.snakeBody.Body[i][1], '#' , nil, g.style.snakeStyle);
	}
	g.Update();
}
func (sb *SnakeBody) Grow(){
	head:= sb.Body[sb.length-1]
	posX:= head[0];
	posY:= head[1];
	if sb.Xspeed == 1 && sb.Yspeed == 0{
		posX++;
	} else if sb.Xspeed == -1 && sb.Yspeed == 0{
		posX--;
	} else if sb.Xspeed == 0 && sb.Yspeed == 1{
		posY--;
	} else if sb.Xspeed == 0 && sb.Yspeed == -1{
		posY++;
	}
	sb.Enqueue([2]int {posX, posY});
	sb.length++;
}
func (g *Game) HasEaten() bool{
	head:= g.snakeBody.Body[g.snakeBody.length-1]
	if head[0] == g.food.X && head[1] == g.food.Y{
		return true;
	}
	return false;
}
func (g *Game) Eat(){
	g.Screen.SetContent(g.food.X, g.food.Y, ' ', nil, g.style.defStyle )
}