package game

type Zyanken int

const (
	Gu Zyanken = iota
	Choki
	Pa
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
