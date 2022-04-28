package mouseevent

// EventType has different calls for each type of event
type EventType int32

// EventType has different calls for each type of event
const (
	Click EventType = iota
	DoubleClick
	TripleClick
	Drag
	Drop
)
