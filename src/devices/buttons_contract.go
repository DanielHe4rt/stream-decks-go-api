package devices

type ButtonKey int

type ButtonContract interface {
	IsPressed() bool

	ButtonPressed() ButtonKey
}
