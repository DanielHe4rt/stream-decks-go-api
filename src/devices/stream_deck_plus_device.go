package devices

import (
	"fmt"
	"github.com/karalabe/hid"
)

type StreamDeckPlus struct {
	DeviceBuffer  hid.Device
	currentBuffer []byte
	buttonsCount  int
	knobsCount    int
	hasKnobs      bool
	hasTouchBar   bool
}

func NewStreamDeckPlus(deviceBuffer hid.Device) *StreamDeckPlus {
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

func (s *StreamDeckPlus) Write(payload []byte) error {
	write, err := s.DeviceBuffer.Write(payload)
	if err != nil {
		return err
	}
	fmt.Println("Bytes Sent: ", write)

	return nil
}

func (s *StreamDeckPlus) SetBrightness(brightness int) error {
	payload := []byte{
		0x03,             // Report ID
		0x08,             // Command or mode identifier (example)
		byte(brightness), // Brightness level (dynamic value)
	}
	//payload = append(payload, make([]byte, 1)...)

	fmt.Println("Payload: ", payload)

	_, err := s.DeviceBuffer.SendFeatureReport(payload)
	if err != nil {
		fmt.Println("fudeu")
		return err
	}

	return nil
}
