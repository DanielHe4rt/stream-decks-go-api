package devices

type DisplayContract interface {
	DisplayEnabled() bool
	HasDisplayInteraction() bool
	DisplayAction() (DisplayResponse, error)
}
type DisplayAction byte

const (
	DisplayTouch DisplayAction = 1
	DisplaySwipe DisplayAction = 3
)

var DisplayActions = map[DisplayAction]string{
	DisplayTouch: "Touch",
	DisplaySwipe: "Swipe",
}

type DisplayState int

const (
	DisplayStateHold          DisplayState = 1
	DisplayStateSwipeForward  DisplayState = 2
	DisplayStateSwipeBackward DisplayState = 3
)

func (d DisplayState) String() string {
	switch d {
	case DisplayStateHold:
		return "Hold"
	case DisplayStateSwipeForward:
		return "Swipe Forward"
	case DisplayStateSwipeBackward:
		return "Swipe Backward"
	default:
		return "Unknown"
	}
}

type DisplayInfo struct {
	Width     int
	RealWidth int
	Height    int
	Index     int
	State     DisplayState
}

type DisplayResponse struct {
	Action       DisplayAction
	Interactions []DisplayInfo
}

func (d DisplayAction) String() string {
	return DisplayActions[d]
}
