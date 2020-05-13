package dom

import (
	"syscall/js"
)

type Element struct {
	Self js.Value
}

func NewElement(tagName string) Element {
	return Element{Document.Call("createElement", tagName)}
}

func (e Element) GetElementsByTagName(tagName string) Element {
	return Element{e.Self.Call("getElementsByTagName", tagName)}
}

func (e Element) ElementsByTagName(tagName string) Element {
	return Element{e.Self.Call("getElementsByTagName", tagName)}
}

func (e Element) SetInnerText(text string) {
	e.Self.Set("innerText", text)
}

func (e Element) SetInnerHTML(text string) {
	e.Self.Set("innerHTML", text)
}

func (e Element) AppendChild(child Element) {
	e.Self.Call("appendChild", child.Self)
}
