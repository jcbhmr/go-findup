package findup

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"syscall/js"
)

var importJS = js.Global().Get("Function").New("s", "return import(s)")

func awaitJS(value js.Value) js.Value {
	ch := make(chan js.Value)
	value.Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
		ch <- args[0]
		return nil
	}))
	return <-ch
}

func stringJS(value js.Value) string {
	if value.Type() == js.TypeString {
		return value.String()
	} else {
		panic(&js.ValueError{Method: "stringJS", Type: value.Type()})
	}
}

//go:embed find-up.min.mjs
var find_up_min_mjs []byte

var moduleJS = awaitJS(importJS.Call("data:text/javascript;base64," + base64.StdEncoding.EncodeToString(find_up_min_mjs)))

func findUp(name any, options *Options) (string, bool) {
	var nameJS js.Value
	switch name := name.(type) {
	case string:
		nameJS = js.ValueOf(name)
	case []string:
		nameAny := make([]any, len(name))
		for i, n := range name {
			nameAny[i] = n
		}
		nameJS = js.ValueOf(nameAny)
	default:
		panic(fmt.Sprintf("invalid name type: %T", name))
	}
	optionsJS := js.ValueOf(options)
	resultJS := awaitJS(moduleJS.Call("findUp", nameJS, optionsJS))
	if resultJS.IsUndefined() {
		return "", false
	} else {
		return stringJS(resultJS), true
	}
}

func findUp2(matcher func(directory string) Match, options *Options) (string, bool) {
	panic("todo")
}
