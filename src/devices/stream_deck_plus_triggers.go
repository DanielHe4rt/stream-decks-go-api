package devices

import (
	"fmt"
)

func (s *StreamDeckPlus) TriggerType() (TriggerType, error) {
	payload := [2]byte(s.currentBuffer[0:2])
	switch payload {
	case [2]byte{1, 0}:
		return ButtonTrigger, nil
	case [2]byte{1, 2}:
		return TouchTrigger, nil
	case [2]byte{1, 3}:
		return KnobTrigger, nil

	default:
		return 0, fmt.Errorf("trigger not implemented: %v", payload)
	}
}
