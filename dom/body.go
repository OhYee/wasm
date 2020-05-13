package dom

var Body = BodyObject{Element{Document.Call("getElementsByTagName", "body").Index(0)}}

type BodyObject struct {
	Element
}
