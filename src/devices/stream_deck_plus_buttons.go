package devices

import (
	"bytes"
)

const (
	ButtonReleased ButtonKey = -1
	ButtonK1       ButtonKey = iota
	ButtonK2
	ButtonK3
	ButtonK4
	ButtonK5
	ButtonK6
	ButtonK7
	ButtonK8
)

var ButtonKeys = map[ButtonKey]string{
	ButtonReleased: "K0 - No button pressed",
	ButtonK1:       "K1 - First Row, First Button (1x0)",
	ButtonK2:       "K2 - First Row, Second Button (1x1)",
	ButtonK3:       "K3 - First Row, Third Button (1x2)",
	ButtonK4:       "K4 - First Row, Fourth Button (1x3)",
	ButtonK5:       "K5 - Second Row, First Button (2x0)",
	ButtonK6:       "K6 - Second Row, Second Button (2x1)",
	ButtonK7:       "K7 - Second Row, Third Button (2x2)",
	ButtonK8:       "K8 - Second Row, Fourth Button (2x3)",
}

func (b ButtonKey) String() string {
	return ButtonKeys[b]
}

// IsPressed
// The button is pressed when one of the byte is 1
// [0, 0, 0, 0, 0, 0, 0, 0] -> No button pressed
// [0, 0, 0, 0, 0, 0, 0, 1] -> Some button pressed
func (s *StreamDeckPlus) IsPressed() bool {
	buttonsBytes := s.currentBuffer[4:]

	pressedButton := []byte{1}

	return bytes.Contains(buttonsBytes, pressedButton)
}

// ButtonPressed
// The button pressed is the one which byte is equal 1
// [0, 0, 0, 0, 0, 0, 0, 0] -> No button pressed
// [0, 0, 0, 0, 0, 0, 0, 1] -> Button 8 pressed
func (s *StreamDeckPlus) ButtonPressed() ButtonKey {
	buttonsBytes := s.currentBuffer[4:]

	pressedButton := []byte{1}

	pressedButtonResponse := bytes.Index(buttonsBytes, pressedButton)

	if pressedButtonResponse == -1 {
		return ButtonReleased
	}

	return ButtonKey(pressedButtonResponse + 1)
}
