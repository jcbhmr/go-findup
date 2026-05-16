//go:build !js

package findup

func findUp(name any, options *Options) (string, bool) {
	var names []string
	switch name := name.(type) {
	case string:
		names = []string{name}
	case []string:
		names = name
	default:
		panic(fmt.Sprintf("invalid name type: %T", name))
	}
	
}