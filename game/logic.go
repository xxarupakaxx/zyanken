package game

import "fmt"

type Game struct {
	started  bool
	finished bool
	te       Zyanken
}

func (g *Game) DecideWinLose(aite Zyanken) bool {
	result := zyankenLogic(g.te, aite)
	if result == Draw {
		ReGame()
	}

	fmt.Println("finished")

	if result == Lose {
		return false
	}

	return true
}

func display(te Zyanken) {
	fmt.Println("")

	fmt.Printf("You: %v\n", ConvertZyankenToStr(te))

	fmt.Println("ーーーーーーーーーーーーーー")
}

func zyankenLogic(watashi, aite Zyanken) Result {
	ans := (watashi - aite + 3) % 3

	return Result(ans)
}
