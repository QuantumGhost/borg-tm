package consts

import (
	"fmt"
	"os"
)

var commitID string
var version string

func CommitID() string {
	return commitID
}

func Version() string {
	return version
}

func PrintVersion() {
	const tmpl = `Borg-TM

Version: %s
CommitID: %s
`
	fmt.Fprintf(os.Stderr, tmpl, version, commitID)
	os.Exit(0)
}
