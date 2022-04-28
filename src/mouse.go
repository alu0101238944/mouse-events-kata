package mouseevent

const timeWindowInMillisecondsForDoubleClick = 500

type state int32

const (
	cleared state = iota
	clicked
	doubleClicked
	tripleClicked
)

// Mouse has an array of listeners
type Mouse struct {
	listeners           []Listener
	lastStateChangeTime uint32
	currentState        state
}

func (mouse *Mouse) PressLeftButton(currentTimeInMilliseconds uint32) {
	/*... implement this method ...*/
}

func (mouse *Mouse) ReleaseLeftButton(currentTimeInMilliseconds uint32) {
	if currentTimeInMilliseconds-mouse.lastStateChangeTime > timeWindowInMillisecondsForDoubleClick {
		mouse.currentState = cleared
	}
	switch mouse.currentState {
	case cleared:
		mouse.notifySubscribers(Click)
		mouse.currentState = clicked
	case clicked:
		mouse.notifySubscribers(DoubleClick)
		mouse.currentState = doubleClicked
	case doubleClicked:
		mouse.notifySubscribers(TripleClick)
		mouse.currentState = tripleClicked
	}
	mouse.lastStateChangeTime = currentTimeInMilliseconds
}

func (mouse *Mouse) Move(from MouseCoordinates, to MouseCoordinates,
	currentTimeInMilliseconds uint32) {
	/*... implement this method ...*/
}

func (mouse *Mouse) Subscribe(listener Listener) {
	mouse.listeners = append(mouse.listeners, listener)
}

func (mouse *Mouse) notifySubscribers(eventType EventType) {
	for _, listener := range mouse.listeners {
		listener.HandleMouseEvent(eventType)
	}
}
