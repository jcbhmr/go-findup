package findup

type findUpStopUniqueSymbol bool

// Return this in a matcher function to stop the search and force FindUp to immediately return nil.
const FindUpStop findUpStopUniqueSymbol = false

// string | typeof(FindUpStop) | nil
type Match any

type Options struct {
	// A directory path where the search halts if no matches are found before reaching this point.
  //
	// Default: Root directory
  StopAt *string
  locatepath.Options
}

// Find a file or directory by walking up parent directories.
//
// Params:
//
// - name: string | []string: The name of the file or directory to find. Can be multiple.
// - options
//
// Returns: The first path found (by respecting the order of names) or nil if none could be found.
func FindUp(name any, options *Options) future[*string] {
  return goFuture(func () (string, error) {
    var options2 Options
    if options != nil {
      options2 = *options
    }
    var cwdRel string
    if options2.Cwd != nil {
      cwdRel = unicornmagic.ToPath(*options2.Cwd)
    }
    directory, err := filepath.Abs(cwdRel)
    if err != nil {
      return 
    }
    // root := ...
  })
}

// Find a file or directory by walking up parent directories.
//
// Params:
//
// - matcher: Called for each directory in the search. Return a path or FindUpStop to stop the search.
//
// Returns: The first path found or nil if none could be found.
func FindUp2(matcher func(directory string) Match, options *Options) future[*string] {
  panic("unimplemented")
}

func FindUpMultiple(name any, options *Options) future[[]string] {}

func FindUpMultiple2(name any, options *Options) future[[]string] {}

var PathExists = pathexists.PathExists
