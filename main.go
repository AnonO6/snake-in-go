package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
)
func main()  {
	screen, err := tcell.NewScreen();

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil{
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite);
	screen.SetStyle(defStyle);
	var w,h = screen.Size();
	snakeBody := SnakeBody{
		Xspeed: 1,
		Yspeed: 0,
		Body: [][2]int{{5,10},},
		length: 1,
	}
	food := Food{
		X: [3]int{6,7,8},
		Y: [3]int{11,12,13},
	}
	style:= Style{
		defStyle: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite),
		snakeStyle: tcell.StyleDefault.Background(tcell.ColorLightGreen).Foreground(tcell.ColorLightGreen),
		foodStyle: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorOrangeRed),
	}
	game := Game{
		Screen:    screen,
		snakeBody: snakeBody,
		food: food,
		width: w,
		height: h,
		style: style,
	}
	go game.Run();

	for{
		switch event:= game.Screen.PollEvent().(type){
		case *tcell.EventResize:
			game.Screen.Sync()
			wn,hn := event.Size()
			game.width= wn;
			game.height= hn;
		case *tcell.EventKey:
			if event.Key() == tcell.KeyCtrlC || event.Key() == tcell.KeyEscape {
				game.Screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyUp {
				game.snakeBody.ChangeDir(0, -1) //In tcell up is -1 and down is +1
			} else if event.Key() == tcell.KeyDown {
				game.snakeBody.ChangeDir(0, 1)
			} else if event.Key() == tcell.KeyLeft {
				game.snakeBody.ChangeDir(-1, 0)
			} else if event.Key() == tcell.KeyRight {
				game.snakeBody.ChangeDir(1, 0)
			}
		}
	}
}