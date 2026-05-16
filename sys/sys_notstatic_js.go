//go:build !static

package sys

import (
	_ "embed"
	"encoding/json"
	"os"
	"sync"
	"syscall/js"

	"go.jcbhmr.com/findup/internal/jsutil"
	"golang.org/x/mod/semver"
)

//go:embed find-up.js
var find_up_js []byte

var nodeProcess = sync.OnceValue(func() js.Value {
	v := js.Global().Get("process")
	if v.IsUndefined() {
		panic("process is undefined")
	}
	return v
})

var nodeModule = sync.OnceValue(func() js.Value {
	return nodeProcess().Call("getBuiltinModule", "module")
})

var version = func() string {
	p := jsutil.AsString(nodeModule().Call("findPackageJSON", "find-up"))
	f, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	var data struct {
		Version string `json:"version"`
	}
	err = dec.Decode(&data)
	if err != nil {
		panic(err)
	}
	v := "v" + data.Version
	if !semver.IsValid(v) {
		panic("invalid version: " + v)
	}
	return v
}()

func init() {
	lowerInclusive := "v8.0.0"
	upperExclusive := "v9.0.0"
	if semver.Compare(version, lowerInclusive) < 0 || semver.Compare(version, upperExclusive) >= 0 {
		panic("unsupported version: " + version)
	}
}

var m = jsutil.Await(jsutil.Import("find-up", nil))
