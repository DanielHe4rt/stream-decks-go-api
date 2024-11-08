package main

import (
	"fmt"
	"github.com/andreykaipov/goobs"
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
		_, err := device.ReadInput()

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
			fmt.Println("Knob Enabled: ", device.KnobEnabled())

			knobResponse, err := device.KnobAction()
			if err != nil {
				fmt.Println("Failed to get knob action:", err)
			}
			fmt.Println("Knob Action: ", knobResponse.Action)
			fmt.Println("Knob Index: ", knobResponse.Index)
			fmt.Println("Knob Value: ", knobResponse.Value)
		} else if trigger == devices.TouchTrigger {
			if !device.DisplayEnabled() {
				fmt.Println("Touch Enabled")
			}
			response, err := device.DisplayAction()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Action: ", response.Action)
			for _, interaction := range response.Interactions {
				fmt.Printf(" --------------------\n")
				fmt.Printf(" Width: %d\n", interaction.Width)
				fmt.Printf(" Height: %d\n", interaction.Height)
				fmt.Printf(" Index: %d\n", interaction.Index)
				fmt.Printf(" Real Width: %d\n", interaction.RealWidth)
				fmt.Printf(" State: %v\n", interaction.State)
				fmt.Printf(" --------------------\n")
			}
		}

	}
}

func main() {

	connectedDevices, err := hid.Enumerate(ElgatoVendor, devices.StreamDeckPlusDevice)

	if err != nil {
		fmt.Printf("Stream Deck not found: %v \n", err)
		return
	}
	deviceInfo := connectedDevices[0]

	hardwareDeviceBuffer, err := deviceInfo.Open()
	if err != nil {
		fmt.Println("Failed to open device:", err)
		return
	}

	device, err := devices.GetDevice(deviceInfo.ProductID, hardwareDeviceBuffer)
	if err != nil {
		// Device Not Supported :x
		log.Fatal(err)
	}
	defer hardwareDeviceBuffer.Close()

	client, err := goobs.New("localhost:1337", goobs.WithPassword("goodpassword"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect()

	version, err := client.General.GetVersion()
	if err != nil {
		panic(err)
	}

	client.Filters.
		fmt.Printf("OBS Studio version: %s\n", version.ObsVersion)
	fmt.Printf("Server protocol version: %s\n", version.ObsWebSocketVersion)
	fmt.Printf("Client protocol version: %s\n", goobs.ProtocolVersion)
	fmt.Printf("Client library version: %s\n", goobs.LibraryVersion)

	fmt.Println("Connected to Stream Deck.")

	go listenForKeyStates(device)

	select {}
}
