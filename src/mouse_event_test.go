package mouseevent

import (
	"testing"
)

type mockListener struct {
	events []EventType
}

func (listener *mockListener) HandleMouseEvent(eventType EventType) {
	listener.events = append(listener.events, eventType)
}

func equalSlices(a, b []EventType) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func sendInstantClick(mouse *Mouse, time Milliseconds) {
	mouse.PressLeftButton(time)
	mouse.ReleaseLeftButton(time)
}

func TestMouse(t *testing.T) {
	t.Run("should call all the listeners with the click event when a click occurs", func(T *testing.T) {
		mouse := Mouse{}
		listener := mockListener{make([]EventType, 0)}
		mouse.Subscribe(&listener)

		sendInstantClick(&mouse, 0)

		actual := listener.events
		expected := []EventType{Click}
		if !equalSlices(actual, expected) {
			t.Errorf("Click didn't happen when it should\nActual: %d | Expected: %d", actual, expected)
		}
	})

	t.Run("should call all the listeners with the double click event when a double click occurs", func(T *testing.T) {
		mouse := Mouse{}
		listener := mockListener{make([]EventType, 0)}
		mouse.Subscribe(&listener)

		sendInstantClick(&mouse, 0)
		sendInstantClick(&mouse, 0)

		actual := listener.events
		expected := []EventType{Click, DoubleClick}
		if !equalSlices(actual, expected) {
			t.Errorf("DoubleClick didn't happen when it should\nActual: %d | Expected: %d", actual, expected)
		}
	})

	t.Run("should call all the listeners with the triple click event when a triple click occurs", func(T *testing.T) {
		mouse := Mouse{}
		listener := mockListener{make([]EventType, 0)}
		mouse.Subscribe(&listener)

		sendInstantClick(&mouse, 0)
		sendInstantClick(&mouse, 0)
		sendInstantClick(&mouse, 0)

		actual := listener.events
		expected := []EventType{Click, DoubleClick, TripleClick}
		if !equalSlices(actual, expected) {
			t.Errorf("TripleClick didn't happen when it should\nActual: %d | Expected: %d", actual, expected)
		}
	})

	t.Run("shouldn't detect two clicks as double click", func(T *testing.T) {
		mouse := Mouse{}
		listener := mockListener{make([]EventType, 0)}
		mouse.Subscribe(&listener)

		sendInstantClick(&mouse, 0)
		sendInstantClick(&mouse, 500)

		actual := listener.events
		expected := []EventType{Click, Click}
		if !equalSlices(actual, expected) {
			t.Errorf("DoubleClick detected when two clicks happened\nActual: %d | Expected: %d", actual, expected)
		}
	})

	t.Run("should call all the listeners with the drag event when a drag occurs", func(T *testing.T) {
		mouse := Mouse{}
		listener := mockListener{make([]EventType, 0)}
		mouse.Subscribe(&listener)

		mouse.PressLeftButton(0)
		mouse.Move(MouseCoordinates{10, 10}, MouseCoordinates{20, 20}, 100)

		actual := listener.events
		expected := []EventType{Click, Drag}
		if !equalSlices(actual, expected) {
			t.Errorf("Drag didn't happen when it should\nActual: %d | Expected: %d", actual, expected)
		}
	})

	t.Run("should call all the listeners with the drop event when a drop occurs", func(T *testing.T) {
		mouse := Mouse{}
		listener := mockListener{make([]EventType, 0)}
		mouse.Subscribe(&listener)

		mouse.PressLeftButton(0)
		mouse.Move(MouseCoordinates{10, 10}, MouseCoordinates{20, 20}, 100)
		mouse.ReleaseLeftButton(200)

		actual := listener.events
		expected := []EventType{Click, Drag, Drop}
		if !equalSlices(actual, expected) {
			t.Errorf("Drop didn't happen when it should\nActual: %d | Expected: %d", actual, expected)
		}
	})
}
