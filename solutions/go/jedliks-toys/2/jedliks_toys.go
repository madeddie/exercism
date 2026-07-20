// Package jedlik does something.
package jedlik

import "fmt"

// Drive updates battery and distance of Car based on speed and batteryDrain
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// DisplayDistance returns the current distance driven by Car
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", car.distance)
}

// DisplayBattery returns the current battery status of Car
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", car.battery)
}

// CanFinish returns a bool indicating of the Car can finish the track by distance
func (car Car) CanFinish(trackDistance int) bool {
	return car.battery*car.speed >= trackDistance*car.batteryDrain
}
