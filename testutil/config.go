package testutil

import "time"

var (
	Once bool

	// Questions for writing struct + method to facilitate break points
	DebugCallIndex int

	DebugTLE = 2 * time.Second

	AssertOutput = true
)
