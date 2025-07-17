package main

import (
	"fmt"
)

// Interface Vehicle yang berfungsi sebagai abstract class
type Vehicle interface {
	SetMaxSpeed(speed int)
	GetMaxSpeed() int
	CalculateFuelConsumption(distance int) float64
	DisplayInfo()
	GetLicensePlate() string
}

// BaseVehicle struct untuk menyimpan properti dasar
type BaseVehicle struct {
	licensePlate string
	maxSpeed     int
}

// Constructor untuk BaseVehicle
func NewBaseVehicle(licensePlate string) BaseVehicle {
	return BaseVehicle{
		licensePlate: licensePlate,
		maxSpeed:     0,
	}
}

// Method untuk get license plate
func (v *BaseVehicle) GetLicensePlate() string {
	return v.licensePlate
}

// Method untuk set max speed
func (v *BaseVehicle) SetMaxSpeed(speed int) {
	v.maxSpeed = speed
}

// Method untuk get max speed
func (v *BaseVehicle) GetMaxSpeed() int {
	return v.maxSpeed
}

// Method konkret untuk display info
func (v *BaseVehicle) DisplayInfo() {
	fmt.Printf("License Plate: %s\n", v.licensePlate)
	fmt.Printf("Max Speed: %d km/h\n", v.maxSpeed)
}

// ==================== CAR CLASS ====================

// Car struct yang mengimplementasikan Vehicle interface
type Car struct {
	BaseVehicle
	fuelEfficiency int // km per liter
}

// Constructor untuk Car
func NewCar(licensePlate string, fuelEfficiency int) *Car {
	return &Car{
		BaseVehicle:    NewBaseVehicle(licensePlate),
		fuelEfficiency: fuelEfficiency,
	}
}

// Implementasi method CalculateFuelConsumption untuk Car
func (c *Car) CalculateFuelConsumption(distance int) float64 {
	if c.fuelEfficiency == 0 {
		return 0
	}
	return float64(distance) / float64(c.fuelEfficiency)
}

// Getter untuk fuelEfficiency
func (c *Car) GetFuelEfficiency() int {
	return c.fuelEfficiency
}

// Setter untuk fuelEfficiency
func (c *Car) SetFuelEfficiency(efficiency int) {
	c.fuelEfficiency = efficiency
}

// Override DisplayInfo untuk Car
func (c *Car) DisplayInfo() {
	fmt.Println("=== CAR INFO ===")
	c.BaseVehicle.DisplayInfo()
	fmt.Printf("Fuel Efficiency: %d km/l\n", c.fuelEfficiency)
	fmt.Printf("Vehicle Type: Car\n")
}

// ==================== TRUCK CLASS ====================

// Truck struct yang mengimplementasikan Vehicle interface
type Truck struct {
	BaseVehicle
	fuelEfficiency int // km per liter
	cargoWeight    int // kg
}

// Constructor untuk Truck
func NewTruck(licensePlate string, fuelEfficiency int, cargoWeight int) *Truck {
	return &Truck{
		BaseVehicle:    NewBaseVehicle(licensePlate),
		fuelEfficiency: fuelEfficiency,
		cargoWeight:    cargoWeight,
	}
}

// Implementasi method CalculateFuelConsumption untuk Truck
func (t *Truck) CalculateFuelConsumption(distance int) float64 {
	if t.fuelEfficiency == 0 {
		return 0
	}
	baseFuel := float64(distance) / float64(t.fuelEfficiency)
	cargoFuel := float64(t.cargoWeight) * 0.05
	return baseFuel + cargoFuel
}

// Getter untuk fuelEfficiency
func (t *Truck) GetFuelEfficiency() int {
	return t.fuelEfficiency
}

// Setter untuk fuelEfficiency
func (t *Truck) SetFuelEfficiency(efficiency int) {
	t.fuelEfficiency = efficiency
}

// Getter untuk cargoWeight
func (t *Truck) GetCargoWeight() int {
	return t.cargoWeight
}

// Setter untuk cargoWeight
func (t *Truck) SetCargoWeight(weight int) {
	t.cargoWeight = weight
}

// Override DisplayInfo untuk Truck
func (t *Truck) DisplayInfo() {
	fmt.Println("=== TRUCK INFO ===")
	t.BaseVehicle.DisplayInfo()
	fmt.Printf("Fuel Efficiency: %d km/l\n", t.fuelEfficiency)
	fmt.Printf("Cargo Weight: %d kg\n", t.cargoWeight)
	fmt.Printf("Vehicle Type: Truck\n")
}

// ==================== HELPER FUNCTIONS ====================

// Function untuk demo fuel consumption calculation
func demonstrateFuelConsumption(vehicle Vehicle, distance int) {
	fmt.Printf("\n--- Fuel Consumption Calculation ---\n")
	fmt.Printf("Distance: %d km\n", distance)
	fuel := vehicle.CalculateFuelConsumption(distance)
	fmt.Printf("Fuel needed: %.2f liters\n", fuel)
}

// Function untuk test polymorphism
func testPolymorphism(vehicles []Vehicle) {
	fmt.Println("\n========== POLYMORPHISM TEST ==========")
	for i, vehicle := range vehicles {
		fmt.Printf("\n--- Vehicle %d ---\n", i+1)
		vehicle.DisplayInfo()
		demonstrateFuelConsumption(vehicle, 100)
	}
}

// ==================== MAIN FUNCTION ====================

func main() {
	fmt.Println("========== VEHICLE HIERARCHY DEMO ==========")
	
	// Membuat objek Car
	car := NewCar("B-1234-CD", 15) // 15 km/l
	car.SetMaxSpeed(180)
	
	// Membuat objek Truck
	truck := NewTruck("B-5678-EF", 8, 2000) // 8 km/l, 2000 kg cargo
	truck.SetMaxSpeed(120)
	
	// Test individual objects
	fmt.Println("\n========== INDIVIDUAL OBJECT TESTS ==========")
	
	// Test Car
	fmt.Println("\n--- Testing Car ---")
	car.DisplayInfo()
	demonstrateFuelConsumption(car, 150)
	
	// Test Truck
	fmt.Println("\n--- Testing Truck ---")
	truck.DisplayInfo()
	demonstrateFuelConsumption(truck, 150)
	
	// Test setters
	fmt.Println("\n--- Testing Setters ---")
	car.SetFuelEfficiency(20)
	truck.SetCargoWeight(3000)
	
	fmt.Printf("Car fuel efficiency updated to: %d km/l\n", car.GetFuelEfficiency())
	fmt.Printf("Truck cargo weight updated to: %d kg\n", truck.GetCargoWeight())
	
	// Test polymorphism
	vehicles := []Vehicle{car, truck}
	testPolymorphism(vehicles)
	
	// Test dengan berbagai jarak
	fmt.Println("\n========== FUEL CONSUMPTION COMPARISON ==========")
	distances := []int{50, 100, 200, 500}
	
	for _, distance := range distances {
		fmt.Printf("\n--- Distance: %d km ---\n", distance)
		carFuel := car.CalculateFuelConsumption(distance)
		truckFuel := truck.CalculateFuelConsumption(distance)
		
		fmt.Printf("Car fuel consumption: %.2f liters\n", carFuel)
		fmt.Printf("Truck fuel consumption: %.2f liters\n", truckFuel)
		fmt.Printf("Difference: %.2f liters\n", truckFuel-carFuel)
	}
	
	// Test edge cases
	fmt.Println("\n========== EDGE CASES TEST ==========")
	
	// Test dengan efisiensi 0
	carZero := NewCar("B-0000-ZZ", 0)
	fmt.Printf("Car with 0 efficiency fuel consumption (100km): %.2f liters\n", 
		carZero.CalculateFuelConsumption(100))
	
	// Test dengan jarak 0
	fmt.Printf("Car fuel consumption for 0 distance: %.2f liters\n", 
		car.CalculateFuelConsumption(0))
	
	fmt.Println("\n========== PROGRAM COMPLETED ==========")
}
