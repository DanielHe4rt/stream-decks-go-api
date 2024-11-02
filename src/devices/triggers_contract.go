package devices

import "fmt"

type TriggerType uint

const (
	ButtonTrigger TriggerType = iota
	KnobTrigger
	TouchTrigger
	FaderTrigger
)

var TriggerTypes = map[TriggerType]string{
	ButtonTrigger: "Button",
	KnobTrigger:   "Knob",
	TouchTrigger:  "Touch",
	FaderTrigger:  "Fader",
}

type TriggerContract interface {
	TriggerType() (TriggerType, error)
}

func (tt TriggerType) String() string {
	return fmt.Sprintf("%v Trigger", TriggerTypes[tt])
}
