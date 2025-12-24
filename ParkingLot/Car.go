package main

import (
	"fmt"
)

type Car struct {
	Plate       string
	Owner       string
	HoursParked int
}

type ParckingLot struct {
	Cars map[string]Car
}

func NewParkingLot() *ParckingLot {
	return &ParckingLot{
		Cars: make(map[string]Car),
	}
}

func (pl *ParckingLot) ParkCar(car Car) error {
	// Plate must not be empty
	if car.Plate == "" {
		return fmt.Errorf("car plate cannot be empty")
	}
	//Plate must be unique
	if _, exists := pl.Cars[car.Plate]; exists {
		return fmt.Errorf("car with plate %s is already parked", car.Plate)
	}
	//Owner must not be empty
	if car.Owner == "" {
		return fmt.Errorf("car owner cannot be empty")
	}

	pl.Cars[car.Plate] = car
	return nil
}

func (pl *ParckingLot) UnparkCar(plate string) error {
	if _, exists := pl.Cars[plate]; !exists {
		return fmt.Errorf("car with plate %s is not parked", plate)
	}
	delete(pl.Cars, plate)
	return nil
}

func (pl *ParckingLot) UpdateHours(plate string, hours int) error {
	// car exists
	if car, exists := pl.Cars[plate]; exists {
		car.HoursParked = hours
		pl.Cars[plate] = car
		return nil
	}
	// hours must be positive
	if hours < 0 {
		return fmt.Errorf("hours parked cannot be negative")
	}
	// return error if car does not exist
	return fmt.Errorf("car with plate %s is not parked", plate)
}

// Returning all parked cars as slice
func (pl *ParckingLot) ListCars() []Car {
	cars := make([]Car, 0, len(pl.Cars))
	for _, car := range pl.Cars {
		cars = append(cars, car)
	}
	return cars
}

// Bills 100 Tenge per hour

func (pl *ParckingLot) Bills(ratePerHour int) map[string]int {
	bills := make(map[string]int)
	for _, car := range pl.Cars {
		bills[car.Plate] = car.HoursParked * ratePerHour
	}
	return bills
}
