package dom

type Span struct {
	Element
}

func NewSpan() Span {
	return Span{NewElement("span")}
}
