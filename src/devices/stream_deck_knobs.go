package devices

import (
	"bytes"
	"fmt"
	"sync"
)

const (
	KnobButtonK1 KnobIndex = iota
	KnobButtonK2
	KnobButtonK3
	KnobButtonK4
)

var KnobButtons = map[KnobIndex]string{
	KnobButtonK1: "K1",
	KnobButtonK2: "K2",
	KnobButtonK3: "K3",
	KnobButtonK4: "K4",
}

func (k KnobIndex) String() string {
	return KnobButtons[k]
}

func (s *StreamDeckPlus) KnobEnabled() bool {
	return s.hasKnobs
}

var knobPool = sync.Pool{
	New: func() interface{} {
		// Return a new KnobIndex initialized to a neutral value (e.g., -1)
		var k KnobIndex = -1
		return &k
	},
}

// KnobAction
// [1 3 5 0 1 1] - Clockwise
// [1 3 5 0 1 255] - CounterClockwise
// [1 3 5 0 0 1] - Pressed
// [1 3 5 0 0 0] - Released
func (s *StreamDeckPlus) KnobAction() (KnobAction, KnobIndex, int, error) {
	if !s.hasKnobs {
		return 0, 0, 0, fmt.Errorf("Knobs not enabled")
	}

	// Check if the knob is pressed
	if s.currentBuffer[4] == 0 {
		knobsBytes := s.currentBuffer[5 : 5+s.knobsCount]
		knobPressed := []byte{1}

		knobPressedIndex := bytes.Index(knobsBytes, knobPressed)

		// Only put into the pool if the knob was pressed (index >= 0)
		if knobPressedIndex != -1 {
			// Use pool object and reset its value
			knobIndex := knobPool.Get().(*KnobIndex)
			*knobIndex = KnobIndex(knobPressedIndex)
			defer knobPool.Put(knobIndex) // Return it to the pool after use

			return KnobPressed, *knobIndex, 0, nil
		}

		// TODO: knob released should return the latest value
		// If no knob pressed, reset value and return it
		knobIndex := knobPool.Get().(*KnobIndex)
		*knobIndex = KnobIndex(knobPressedIndex)
		defer knobPool.Put(knobIndex) // Put it back in the pool

		return KnobReleased, *knobIndex, 0, nil
	}

	// Check if the knob is rotated
	if s.currentBuffer[4] == 1 {
		validKnobs := s.currentBuffer[5 : 5+s.knobsCount]
		// get the index which is different from 0
		for i, knob := range validKnobs {
			if knob != 0 {
				if knob >= 1 && knob <= 50 {
					// Knob value = 1 to 50
					return KnobClockWise, KnobIndex(i), int(knob), nil
				}
				if knob >= 200 {
					// knob value = 256 - knobvalue
					knobValue := 256 - int(knob)

					return KnobCounterClockWise, KnobIndex(i), knobValue, nil
				}
			}
		}

		return 0, 0, 0, fmt.Errorf("Knob action not implemented")
	}

	return 0, 0, 0, fmt.Errorf("Knob action not implemented")
}

func (s *StreamDeckPlus) KnobInteractedIndex() int {
	return 1
}
