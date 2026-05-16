package findup

import "fmt"

type findUpStopType struct{}

// Return this in a matcher function to stop the search and force [FindUp] to immediately return nil.
var FindUpStop = findUpStopType{}

// string | typeof(FindUpStop) | nil
type Match any

type Options struct {
	// A directory path where the search halts if no matches are found before reaching this point.
	//
	// Default: Root directory
	StopAt *string
	// The type of path to match
	//
	// Default: "file"
	Type *OptionsType
	// The number of matches to limit the search to.
	//
	// Default: math.MaxInt
	Limit *int
	// The current working directory.
	//
	// Default: os.Getwd()
	Cwd *string
	// Allow symbolic links to match if they point to the requested path type.
	//
	// Default: true
	AllowSymlinks *bool
}

type OptionsType string

const (
	OptionsTypeFile      OptionsType = "file"
	OptionsTypeDirectory OptionsType = "directory"
	OptionsTypeBoth      OptionsType = "both"
)

func (t OptionsType) check() {
	switch t {
	case OptionsTypeFile, OptionsTypeDirectory, OptionsTypeBoth:
	default:
		panic(fmt.Sprintf("invalid options type: %v", t))
	}
}

func FindUp(name any, options *Options) (string, bool) {
	return findUp(name, options)
}

func FindUp2(matcher func(directory string) Match, options *Options) (string, bool) {
	panic("todo")
}

func FindUpMultiple(name any, options *Options) ([]string, bool) {
	panic("todo")
}

func FindUpMultiple2(name any, options *Options) ([]string, bool) {
	panic("todo")
}
