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
}

type DeckDevice interface {
	DeviceName() string
	TriggerContract
	DeckBuffer
	ButtonContract
	KnobContract
}

func GetDevice(device *hid.Device) (DeckDevice, error) {
	switch device.ProductID {
	case StreamDeckPlusDevice:
		return NewStreamDeckPlus(device), nil
	default:
		return nil, fmt.Errorf("device not suuported")

	}
}
