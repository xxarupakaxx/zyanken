package game

type Game struct {
	started  bool
	finished bool
	te       Zyanken
}




func zyankenLogic(watashi, aite Zyanken) Result {
	ans := (watashi - aite + 3) % 3

	return Result(ans)
}
