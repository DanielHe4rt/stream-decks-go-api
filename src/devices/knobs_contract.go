package devices

type KnobAction int

type KnobIndex int

const (
	KnobClockWise        KnobAction = KnobAction(0)
	KnobCounterClockWise KnobAction = KnobAction(1)
	KnobPressed          KnobAction = KnobAction(2)
	KnobReleased         KnobAction = KnobAction(3)
)

var KnobActions = map[KnobAction]string{
	KnobClockWise:        "ClockWise",
	KnobCounterClockWise: "CounterClockWise",
	KnobPressed:          "Pressed",
	KnobReleased:         "Released",
}

func (k KnobAction) String() string {
	return KnobActions[k]
}

type KnobContract interface {
	KnobEnabled() bool
	KnobAction() (KnobAction, KnobIndex, int, error)
	KnobInteractedIndex() int
}
