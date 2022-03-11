package game

import "fmt"

type Game struct {
	started  bool
	finished bool
}

func (g *Game) Display(te Zyanken) {
	fmt.Println("")
	fmt.Printf("Your Te: %s\n", ConvertZyankenToStr(te))
	fmt.Println("ーーーーーーーーーーーーーー")

}
