//go:build static

package sys

import (
	_ "embed"
	"encoding/base64"

	"go.jcbhmr.com/findup/internal/jsutil"
)

//go:embed find-up.js
var find_up_js []byte

var m = jsutil.Await(jsutil.Import("data:text/javascript;base64,"+base64.StdEncoding.EncodeToString(find_up_js), nil))
