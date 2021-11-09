package elon

import "fmt"

// Drive doc string
func (c *Car) Drive() {
	if c.battery < c.batteryDrain {
		return
	}
	c.distance += c.speed
	c.battery = c.battery - c.batteryDrain
}

// CanFinish doc string
func (c Car) CanFinish(trackDistance int) bool {
	iter := trackDistance / c.speed

	return (iter * c.batteryDrain) <= c.battery
}

// DisplayDistance doc string
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery doc string
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}
