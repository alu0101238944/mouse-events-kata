package mouseevent

const timeWindowInMillisecondsForDoubleClick = 500

// Mouse has an array of listeners
type Mouse struct {
	listeners []Listener
}

func (mouse *Mouse) PressLeftButton(currentTimeInMilliseconds uint32) {
	/*... implement this method ...*/
}

func (mouse *Mouse) ReleaseLeftButton(currentTimeInMilliseconds uint32) {
	/*... implement this method ...*/
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
