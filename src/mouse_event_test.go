package mouseevent

import "testing"

type mockListener struct {
	lastEvent EventType
}

func (listener *mockListener) HandleMouseEvent(eventType EventType) {
	listener.lastEvent = eventType
}

func TestMouse(t *testing.T) {
	t.Run("should call all the listeners with the click event when a click occurs", func(T *testing.T) {
		mouse := Mouse{}
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
		mouse := Mouse{}
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

	t.Run("should call all the listeners with the triple click event when a triple click occurs", func(T *testing.T) {
		mouse := Mouse{}
		listener := mockListener{Drop}
		mouse.Subscribe(&listener)

		mouse.PressLeftButton(0)
		mouse.ReleaseLeftButton(100)

		mouse.PressLeftButton(200)
		mouse.ReleaseLeftButton(300)

		mouse.PressLeftButton(400)
		mouse.ReleaseLeftButton(500)

		actual := listener.lastEvent
		if actual != TripleClick {
			t.Errorf("TripleClick didn't happen when it should\nActual: %d | Expected: %d", actual, TripleClick)
		}
	})

	t.Run("shouldn't detect two clicks as double click", func(T *testing.T) {
		mouse := Mouse{}
		listener := mockListener{DoubleClick}
		mouse.Subscribe(&listener)

		mouse.PressLeftButton(0)
		mouse.ReleaseLeftButton(100)

		mouse.PressLeftButton(600)
		mouse.ReleaseLeftButton(700)

		actual := listener.lastEvent
		if actual == DoubleClick {
			t.Errorf("DoubleClick detected when two clicks happened\nActual: %d | Expected: %d", actual, DoubleClick)
		}
	})

	t.Run("should call all the listeners with the drag event when a drag occurs", func(T *testing.T) {
		mouse := Mouse{}
		listener := mockListener{Drop}
		mouse.Subscribe(&listener)

		mouse.PressLeftButton(0)
		mouse.Move(MouseCoordinates{10, 10}, MouseCoordinates{20, 20}, 100)

		actual := listener.lastEvent
		if actual != Drag {
			t.Errorf("Drag didn't happen when it should\nActual: %d | Expected: %d", actual, Drag)
		}
	})

	t.Run("should call all the listeners with the drop event when a drop occurs", func(T *testing.T) {
		mouse := Mouse{}
		listener := mockListener{Drop}
		mouse.Subscribe(&listener)

		mouse.PressLeftButton(0)
		mouse.Move(MouseCoordinates{10, 10}, MouseCoordinates{20, 20}, 100)
		mouse.ReleaseLeftButton(200)

		actual := listener.lastEvent
		if actual != Drop {
			t.Errorf("Drop didn't happen when it should\nActual: %d | Expected: %d", actual, Drop)
		}
	})
}
