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

func numberToClickEvent(number uint8) EventType {
	switch number {
	case 1:
		return Click
	case 2:
		return DoubleClick
	case 3:
		return TripleClick
	}
	panic("numberToClickEvent can only be called with numbers between 1 and 3")
}
