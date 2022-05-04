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
	listeners          []Listener
	lasClickChangeTime uint32
	currentState       state
	pressed            bool
	moved              bool
}

func (mouse *Mouse) PressLeftButton(currentTimeInMilliseconds uint32) {
	if currentTimeInMilliseconds-mouse.lasClickChangeTime > timeWindowInMillisecondsForDoubleClick {
		mouse.currentState = cleared
	} else {
		mouse.lasClickChangeTime = currentTimeInMilliseconds
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
	}
	mouse.pressed = true
}

func (mouse *Mouse) ReleaseLeftButton(currentTimeInMilliseconds uint32) {
	mouse.pressed = false
	if mouse.moved {
		mouse.notifySubscribers(Drop)
		mouse.moved = false
	}
}

func (mouse *Mouse) Move(from MouseCoordinates, to MouseCoordinates,
	currentTimeInMilliseconds uint32) {
	if mouse.pressed {
		mouse.notifySubscribers(Drag)
		mouse.moved = true
	}
}

func (mouse *Mouse) Subscribe(listener Listener) {
	mouse.listeners = append(mouse.listeners, listener)
}

func (mouse *Mouse) notifySubscribers(eventType EventType) {
	for _, listener := range mouse.listeners {
		listener.HandleMouseEvent(eventType)
	}
}
