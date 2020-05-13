package dom

import (
	"syscall/js"
)

var Global = js.Global()
var Document = Global.Get("document")

