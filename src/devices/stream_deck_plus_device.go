package devices

import (
	"github.com/karalabe/hid"
)

type StreamDeckPlus struct {
	DeviceBuffer  *hid.Device
	currentBuffer []byte
	buttonsCount  int
	knobsCount    int
	hasKnobs      bool
	hasTouchBar   bool
}

func NewStreamDeckPlus(deviceBuffer *hid.Device) *StreamDeckPlus {
	return &StreamDeckPlus{
		DeviceBuffer: deviceBuffer,
		buttonsCount: 8,
		knobsCount:   4,
		hasKnobs:     true,
		hasTouchBar:  true,
	}
}

func (s *StreamDeckPlus) ReadInput() (bytes []byte, err error) {
	buffer := make([]byte, 32)
	_, err = s.DeviceBuffer.Read(buffer)
	if err != nil {
		return nil, err
	}

	s.currentBuffer = buffer
	return buffer, nil
}

func (s *StreamDeckPlus) DeviceName() string {
	return "StreamDeck+"
}
