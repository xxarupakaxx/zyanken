package game

type Zyanken int
type Result int

const (
	Gu Zyanken = iota
	Choki
	Pa
	None
)

const (
	Draw Result = iota
	Lose
	Win
)

func ConvertZyankenToStr(zyanken Zyanken) string {
	switch zyanken {
	case Gu:
		return "π"
	case Choki:
		return "β"
	case Pa:
		return "π"
	case None:
		return "δ½γιΈζγ§γγ¦γγΎγγ"
	}

	return ""
}

func ConvertResultToStr(result Result) string {
	switch result {
	case Draw:
		return "εΌγεγ"
	case Lose:
		return "θ² γ"
	case Win:
		return "εγ‘"
	}

	return "εΌγεγ"
}
