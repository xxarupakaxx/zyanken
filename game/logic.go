package game

import "fmt"

type Game struct {
	started  bool
	finished bool
	me       Zyanken
}

func NewGame() *Game {
	return &Game{}
}

// DecideWinLose あいこだったらfalseで勝敗が決定したらtrue
func (g *Game) DecideWinLose(aite Zyanken) bool {
	g.Display(g.me)
	result := zyankenLogic(g.me, aite)
	if result == Draw {
		return false
	}

	fmt.Println("finished")
	return true
}

func (g *Game) Display(te Zyanken) {
	fmt.Println("")

	fmt.Printf("You: %v\n", ConvertZyankenToStr(te))

	fmt.Println("ーーーーーーーーーーーーーー")
}

func zyankenLogic(me, aite Zyanken) Result {
	ans := (me - aite + 3) % 3

	return Result(ans)
}

func (g *Game) Winner(aite Zyanken) Result {
	result := zyankenLogic(g.me, aite)
	if result == Win {
		return Win
	}

	return Lose
}
