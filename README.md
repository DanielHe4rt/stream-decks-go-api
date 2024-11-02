# KaDeck Project

This is a project to support different types of StreamDeck focused devices to develop plugins for them.

## Supported Devices

- [x] StreamDeck Plus
- [ ] StreamDeck Mini
- [ ] StreamDeck XL

## Architecture

Depending on the device, you will have contracts and modules to implement and use it.

Since I'm using the StreamDeck+ as my first (and only) device, I'm focusing on it first.

Under the contracts we have:

- `triggers_contract.go`: This contract is responsible for the triggers that the device will have. It's the main
  contract that will be used to interact with the device.
- `buttons_contract.go`: This contract is responsible for the buttons that the device will have by parsing and reusing.
  This will be the most common among StreamDecks.
- `knobs_contract.go`: This contract is responsible for the knobs that the device will have. This is a specific contract
  for the StreamDeck+ and other similar devices that have knobs.

If you want to implement a new device, you should follow the steps below:

```go
// In `new_device.go`
package devices

import (
    "github.com/karalabe/hid"
)

type NewDevice struct {
    device *hid.Device
    currentBuffer []byte
}

func NewNewDevice(device *hid.Device) *NewDevice {
    return &NewDevice{device: device}
}

// Implement DeckDevice, ButtonContract, KnobContract interfaces
func (d *NewDevice) DeviceName() string {
    return "New Device"
}

func (d *NewDevice) ReadInput() ([]byte, error) {
    buffer := make([]byte, 64)
    _, err := d.device.Read(buffer)
    if err != nil {
        return nil, err
    }
    d.currentBuffer = buffer
    return buffer, nil
}

func (d *NewDevice) IsPressed() bool {
    // Implement button press logic
    return false
}

func (d *NewDevice) ButtonPressed() ButtonKey {
    // Implement button press logic
    return ButtonReleased
}

func (d *NewDevice) KnobEnabled() bool {
    // Implement knob enabled logic
    return false
}

func (d *NewDevice) KnobAction() (KnobAction, KnobIndex, int, error) {
    // Implement knob action logic
    return KnobClockWise, 0, 0, nil
}

func (d *NewDevice) KnobInteractedIndex() int {
    // Implement knob interacted index logic
    return 0
}

```

It uses the `github.com/karalabe/hid` package for HID (Human Interface Device) communication.