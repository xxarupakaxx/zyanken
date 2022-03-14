package game

type Zyanken int
type Result int

const (
	Gu Zyanken = iota
	Choki
	Pa
)

const (
	Draw Result = iota
	Lose
	Win
)

func ConvertZyankenToStr(zyanken Zyanken) string {
	switch zyanken {
	case Gu:
		return "👊"
	case Choki:
		return "✌"
	case Pa:
		return "🖐"
	}

	return ""
}

func ConvertResultToStr(result Result) string {
	switch result {
	case Draw:
		return "引き分け"
	case Lose:
		return "負け"
	case Win:
		return "勝ち"
	}

	return ""
}
