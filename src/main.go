package main

import (
	"fmt"
	"github.com/danielhe4rt/go-fodase/src/devices"
	"github.com/karalabe/hid"
	"log"
)

type VendorType = uint16

const (
	ElgatoVendor VendorType = VendorType(4057)
)

func listenForKeyStates(device devices.DeckDevice) {

	for {
		buf, err := device.ReadInput()

		if err != nil {
			fmt.Println("Failed to read from device:", err)
			return
		}

		trigger, err := device.TriggerType()
		if err != nil {
			fmt.Println("Failed to get trigger type:", err)
		}

		if trigger == devices.ButtonTrigger {
			fmt.Println(device.ButtonPressed())
			fmt.Println(device.IsPressed())
		} else if trigger == devices.KnobTrigger {
			fmt.Println("----------------")
			fmt.Println(buf)
			fmt.Println("Knob Enabled: ", device.KnobEnabled())

			knobAction, knobIndex, knobValue, err := device.KnobAction()
			if err != nil {
				fmt.Println("Failed to get knob action:", err)
			}
			fmt.Println("Knob Action: ", knobAction)
			fmt.Println("Knob Index: ", knobIndex)
			fmt.Println("Knob Value: ", knobValue)
		}

	}
}

func main() {

	connectedDevices := hid.Enumerate(ElgatoVendor, devices.StreamDeckPlusDevice)

	if len(connectedDevices) == 0 {
		fmt.Println("Stream Deck not found.")
		return
	}
	deviceInfo := connectedDevices[0]

	hardwareDeviceBuffer, err := deviceInfo.Open()
	if err != nil {
		fmt.Println("Failed to open device:", err)
		return
	}

	device, err := devices.GetDevice(hardwareDeviceBuffer)
	if err != nil {
		// Device Not Supported :x
		log.Fatal(err)
	}

	defer hardwareDeviceBuffer.Close()

	fmt.Println("Connected to Stream Deck.")

	go listenForKeyStates(device)

	select {}
}
