package dom

import (
	"syscall/js"
)

type Input struct {
	Element
}

func NewInput() Input {
	return Input{NewElement("input")}
}

func (e Input) Value() string {
	return e.Self.Get("value").String()
}

func (e Input) SetOnInput(f func(this Input, args []js.Value) interface{}) {
	e.Self.Set("oninput", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			return f(Input{Element{this}}, args)
		},
	))
}
