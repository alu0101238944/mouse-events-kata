package mouseevent

const timeWindowForClicks = 500

// Milliseconds unit
type Milliseconds int32

// Mouse has an array of listeners
type Mouse struct {
	listeners             []Listener
	lastClick             Milliseconds
	clicksInCurrentWindow uint8
	pressed               bool
	moved                 bool
}

func (mouse *Mouse) PressLeftButton(currentTime Milliseconds) {
	if currentTime-mouse.lastClick > timeWindowForClicks {
		mouse.clicksInCurrentWindow = 0
	}
	mouse.lastClick = currentTime

	eventToNotify := numberToClickEvent(mouse.clicksInCurrentWindow + 1)
	mouse.notifySubscribers(eventToNotify)

	mouse.clicksInCurrentWindow = (mouse.clicksInCurrentWindow + 1) % 3
	mouse.pressed = true
}

func (mouse *Mouse) ReleaseLeftButton(currentTime Milliseconds) {
	mouse.pressed = false
	if mouse.moved {
		mouse.notifySubscribers(Drop)
		mouse.moved = false
	}
}

func (mouse *Mouse) Move(from MouseCoordinates, to MouseCoordinates,
	currentTime Milliseconds) {
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
