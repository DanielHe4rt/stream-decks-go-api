package devices

import (
	"fmt"
	"github.com/karalabe/hid"
)

type ProductType = uint16

const (
	StreamDeckPlusDevice ProductType = ProductType(132)
)

type DeckBuffer interface {
	ReadInput() ([]byte, error)
	Write([]byte) error
}

type DeckDevice interface {
	DeviceName() string
	SetBrightness(brightness int) error
	TriggerContract
	DeckBuffer
	ButtonContract
	KnobContract
	DisplayContract
}

func GetDevice(productId uint16, device hid.Device) (DeckDevice, error) {
	switch productId {
	case StreamDeckPlusDevice:
		return NewStreamDeckPlus(device), nil
	default:
		return nil, fmt.Errorf("device not suuported")

	}
}
