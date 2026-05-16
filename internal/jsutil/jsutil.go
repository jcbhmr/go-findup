//go:build js

package jsutil

import (
	"sync"
	"syscall/js"
)

var promise = sync.OnceValue(func() js.Value {
	return js.Global().Get("Promise")
})

func Await(v js.Value) js.Value {
	p := promise().Call("resolve", v)
	type Result struct {
		Value js.Value
		Err   error
	}
	ch := make(chan Result)
	onResolve := js.FuncOf(func(this js.Value, args []js.Value) any {
		ch <- Result{Value: args[0]}
		return nil
	})
	defer onResolve.Release()
	onReject := js.FuncOf(func(this js.Value, args []js.Value) any {
		ch <- Result{Err: js.Error{Value: args[0]}}
		return nil
	})
	p.Call("then", onResolve, onReject)
	result := <-ch
	if result.Err != nil {
		panic(result.Err)
	}
	return result.Value
}

var function = sync.OnceValue(func() js.Value {
	return js.Global().Get("Function")
})

var importFunc = sync.OnceValue(func() js.Value {
	return function().New("s", "o", "return import(s,o)")
})

func Import(specifier string, options map[string]any) js.Value {
	var optionsJS js.Value
	if options == nil {
		optionsJS = js.Undefined()
	} else {
		optionsJS = js.ValueOf(options)
	}
	return importFunc().Invoke(specifier, optionsJS)
}

func AsString(v js.Value) string {
	if v.Type() != js.TypeString {
		panic(&js.ValueError{
			Method: "Value.String",
			Type:   v.Type(),
		})
	}
	return v.String()
}
