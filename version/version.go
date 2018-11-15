package version

import (
	"fmt"
)

// Version of app
var Version = ""

// CommitHash of app at build time
var CommitHash = ""

// Branch of  of app at build time
var Branch = ""

// BuildTime is timestamp of app compilation
var BuildTime = ""

// String returns full version of this app
func String() string {
	return fmt.Sprintf(
		"%s built at %s from commit %s@%s",
		Version, BuildTime, CommitHash, Branch,
	)
}
