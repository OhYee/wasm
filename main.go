package main

import (
	"fmt"
	"syscall/js"

	"github.com/OhYee/rainbow/errors"
	"github.com/OhYee/wasm/dom"
	wasm "github.com/OhYee/wasm/package"
)

//go:generate bash -c "GOARCH=wasm GOOS=js go build -o main.wasm main.go"
func sayHi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

func main() {
	pkg := wasm.NewPackage("demo")
	pkg.ExportFunction("sayHi", func(this js.Value, args []js.Value) (ret interface{}, err error) {
		if len(args) >= 1 && args[0].Type() == js.TypeString {
			name := args[0].String()
			return sayHi(name), nil
		}
		err = errors.New("Call failed: sayHi(name:string):string need a string argument")
		return
	})

	span := dom.NewSpan()
	span.SetInnerText("input some thing")

	input := dom.NewInput()
	input.SetOnInput(
		func(this dom.Input, args []js.Value) interface{} {
			span.SetInnerText(this.Value())
			return nil
		},
	)

	dom.Body.AppendChild(input.Element)
	dom.Body.AppendChild(span.Element)

	pkg.Run()
}
