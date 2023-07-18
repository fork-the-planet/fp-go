package exec

import (
	T "github.com/ibm/fp-go/tuple"
)

type (
	// command output
	CommandOutput = T.Tuple2[[]byte, []byte]
)

var (
	StdOut = T.First[[]byte, []byte]
	StdErr = T.Second[[]byte, []byte]
)
