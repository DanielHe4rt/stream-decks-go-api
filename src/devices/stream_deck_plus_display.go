package devices

import (
	"bytes"
	"fmt"
)

func (s *StreamDeckPlus) DisplayEnabled() bool {
	return s.hasTouchBar
}

func (s *StreamDeckPlus) HasDisplayInteraction() bool {
	interactionPayload := []byte{1, 2, 14}

	return bytes.Equal(interactionPayload, s.currentBuffer[0:3])
}

func (s *StreamDeckPlus) DisplayAction() (DisplayResponse, error) {
	if !s.hasTouchBar {
		return DisplayResponse{
			Action:       0,
			Interactions: nil,
		}, fmt.Errorf("touch not enabled")
	}

	switch DisplayAction(s.currentBuffer[4]) {
	case DisplayTouch:
		return s.touchAction()
	case DisplaySwipe:
		return s.touchSwipe()
	default:
		return DisplayResponse{
			Action:       0,
			Interactions: nil,
		}, fmt.Errorf("touch action not implemented")
	}

}

/*
Payload Structure:
[1 1 18 0 72]

Index | Description         | Value
-----------------------------------
0     | Action              | 1
1     | Unknown             | 1
2     | Width               | 18
3     | Width Multiplier    | 0
4     | Height              | 72
*/

func (s *StreamDeckPlus) touchSwipe() (DisplayResponse, error) {
	if !s.hasTouchBar {
		return DisplayResponse{
			Action:       0,
			Interactions: nil,
		}, fmt.Errorf("touch not enabled")
	}

	var interactions []DisplayInfo

	fromSwipe := s.currentBuffer[6:10]
	toSwipe := s.currentBuffer[10:]

	fromRealWidth, fromSection := calculateDisplayWidth(int(fromSwipe[0]), int(fromSwipe[1]))
	toRealWidth, toSection := calculateDisplayWidth(int(toSwipe[0]), int(toSwipe[1]))

	fromState, toState := calculateSwipeState(fromRealWidth, toRealWidth)

	interactions = append(interactions, DisplayInfo{
		Width:     int(fromSwipe[0]),
		RealWidth: fromRealWidth,
		Height:    int(fromSwipe[2]),
		Index:     fromSection,
		State:     fromState,
	})

	interactions = append(interactions, DisplayInfo{
		Width:     int(toSwipe[0]),
		RealWidth: toRealWidth,
		Height:    int(toSwipe[2]),
		Index:     toSection,
		State:     toState,
	})

	return DisplayResponse{
		Action:       DisplaySwipe,
		Interactions: interactions,
	}, nil

}

func calculateSwipeState(fromWidth int, toWidth int) (DisplayState, DisplayState) {
	if fromWidth < toWidth {
		return DisplayStateSwipeForward, DisplayStateSwipeBackward
	}

	return DisplayStateSwipeBackward, DisplayStateSwipeForward
}

func (s *StreamDeckPlus) touchAction() (DisplayResponse, error) {
	if !s.hasTouchBar {
		return DisplayResponse{
			Action:       0,
			Interactions: nil,
		}, fmt.Errorf("touch not enabled")
	}

	var interactions []DisplayInfo

	currentPayload := s.currentBuffer[4:]

	realWidth, section := calculateDisplayWidth(int(currentPayload[2]), int(currentPayload[3]))
	interactions = append(interactions, DisplayInfo{
		Width:     int(currentPayload[2]),
		RealWidth: realWidth,
		Height:    int(currentPayload[4]),
		Index:     section,
		State:     DisplayStateHold,
	})

	return DisplayResponse{Action: DisplayTouch, Interactions: interactions}, nil

}

func calculateDisplayWidth(width, multiplier int) (int, int) {
	if multiplier == 0 {
		return width, 0
	}

	multiplier = multiplier * 256
	realWidth := width + multiplier
	section := realWidth / 200

	return realWidth, section
}
