package mouseevent

import "testing"

type mockListener struct {
	lastEvent EventType
}

func (listener *mockListener) HandleMouseEvent(eventType EventType) {
	listener.lastEvent = eventType
}

func TestMouse(t *testing.T) {
	mouse := Mouse{}

	t.Run("should call all the listeners with the click event when a click occurs", func(T *testing.T) {
		listener := mockListener{Drop}
		mouse.Subscribe(&listener)

		mouse.PressLeftButton(0)
		mouse.ReleaseLeftButton(100)

		actual := listener.lastEvent
		if actual != Click {
			t.Errorf("Click didn't happen when it should\nActual: %d | Expected: %d", actual, Click)
		}
	})

	t.Run("should call all the listeners with the double click event when a double click occurs", func(T *testing.T) {
		listener := mockListener{Drop}
		mouse.Subscribe(&listener)

		mouse.PressLeftButton(0)
		mouse.ReleaseLeftButton(100)

		mouse.PressLeftButton(200)
		mouse.ReleaseLeftButton(300)

		actual := listener.lastEvent
		if actual != DoubleClick {
			t.Errorf("DoubleClick didn't happen when it should\nActual: %d | Expected: %d", actual, DoubleClick)
		}
	})
}
