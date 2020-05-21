package wasm

import (
	"syscall/js"

	"github.com/OhYee/rainbow/errors"
)

type jsFunction = func(this js.Value, args []js.Value) (interface{}, error)

// ReturnObject
type ReturnObject struct {
	Success bool
	Return  interface{}
}

// Map transfer ReturnObject to map[string]interface{} for returning to js
func (obj *ReturnObject) Map() map[string]interface{} {
	return map[string]interface{}{
		"success": obj.Success,
		"return":  obj.Return,
	}
}

// Package for WASM
type Package struct {
	name    string
	pkg     map[string]interface{}
	channel chan bool
}

// NewPackage initial a WASM package
func NewPackage(name string) *Package {
	return &Package{
		name:    name,
		pkg:     make(map[string]interface{}),
		channel: make(chan bool),
	}
}

// Exit the package
func (pkg *Package) Exit(this js.Value, args []js.Value) interface{} {
	pkg.channel <- true
	return nil
}

func (pkg *Package) wrapperFunction(name string, f jsFunction) func(js.Value, []js.Value) *ReturnObject {
	return func(this js.Value, args []js.Value) (obj *ReturnObject) {
		defer func() {
			if err := recover(); err != nil {
				obj = &ReturnObject{
					Success: false,
					Return:  errors.ShowStack(errors.New("%+v", err)),
				}
			}
		}()

		ret, err := f(this, args)
		if err != nil {
			return &ReturnObject{
				Success: false,
				Return:  errors.ShowStack(errors.NewErr(err)),
			}
		}
			}
		}
		return &ReturnObject{
			Success: true,
			Return:  ret,
		}
	}
}

// ExportFunction export the function to the global.{package name}.exports.{name}
func (pkg *Package) ExportFunction(name string, f jsFunction) {
	pkg.pkg[name] = js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			return pkg.wrapperFunction(name, f)(this, args).Map()
		},
	)
}

// ExportVar export the variable to global.{package name}.exports.{name}
func (pkg *Package) ExportVar(name string, variable js.Value) {
	pkg.pkg[name] = variable
}

// Run the package
func (pkg *Package) Run() {
	js.Global().Set(pkg.name, map[string]interface{}{
		"exit":    js.FuncOf(pkg.Exit),
		"exports": pkg.pkg,
	})
	<-pkg.channel
}
