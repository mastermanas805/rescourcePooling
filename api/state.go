package api

type State int

// Declare related constants for each direction starting with index 1
const (
	Queued     State = iota + 1 // EnumIndex = 1
	Processing                  // EnumIndex = 2
	Completed                   // EnumIndex = 3
)
