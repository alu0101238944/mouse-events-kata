package mouseevent

// Listener listens for MouseEvents and reacts to them
type Listener interface {
	HandleMouseEvent(eventType EventType)
}
