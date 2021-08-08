package parkinglot

import (
	"log"
)

/*
Following is the skeleton code for our parking lot system:

Enums and Constants: Here are the required enums, data types, and constants:
*/

// VehicleType
const (
	CAR = iota
	TRUCK
	ELECTRIC
	VAN
	MOTORBIKE
)

// ParkingSpottype
const (
	HANDICAPPED_SPOT = iota
	COMPACT_SPOT
	LARGE_SPOT
	MOTORBIKE_SPOT
	ELECTRIC_SPOT
)

// AccountStatus
const (
	ACTIVE = iota
	BLOCKED
	BANNED
	COMPRIMISED
	ARCHIVED
	UNKOWN
)

// ParkingTicketStatus
const (
	ACTIVE_Ticket = iota
	PAID_Ticket
	LOST_Ticket
)

type Address struct {
	StreetAddress string
	City          string
	State         string
	ZipCode       string
	Country       string
}

type Person struct {
	Name    string
	Address Address
	Email   string
	Phone   string
}

/*
Account, Admin, and ParkingAttendant: These classes represent various people that interact with our system:
*/
// Account
type Account struct {
	userName string
	password string
	status   string
	Person
}

type Accounter interface {
	ResetPassword() bool
}

type Admin struct {
	Account
}

type AdminAccounter interface {
	Accounter

	AddParkingFloor(floor ParkingFloor) bool
	AddParkingSpot(floorName, parkingSpot string) bool
}

type AdminDisplayer interface {
	AddParkingDisplayBoard(floorName string, displayBoard ParkingDisplayBoard) bool
	AddEntracePanel(entrancePanel EntrancePanel) bool
	AddExitPanel(exitPanel ExitPanel) bool
}

type ParkingAttendant struct {
	Account
}

type ParkingAttendantHelper interface {
	Accounter
	ProcessTicket(ticketNumber string) bool
}

/*
ParkingSpot: Here is the definition of ParkingSpot and all of its children classes:
*/

type ParkingSpot struct {
	number          string
	free            bool
	vehicle         VehicleManger
	parkingSpotType int
}

func (spot ParkingSpot) Number() string {
	return spot.number
}

func (spot ParkingSpot) Type() int {
	return spot.parkingSpotType
}

// SetSpotType
func (spot ParkingSpot) SetSpotType(spotType int) bool {
	switch spotType {
	case HANDICAPPED_SPOT:
	case COMPACT_SPOT:
	case LARGE_SPOT:
	case MOTORBIKE_SPOT:
	case ELECTRIC_SPOT:
		spot.parkingSpotType = spotType
		break
	default:
		return false
	}
	return true
}

func (spot ParkingSpot) IsFree() bool { return spot.free }

func (spot ParkingSpot) AssignVehicle(vehicle VehicleManger) bool {
	spot.vehicle = vehicle
	spot.free = false
	return true
}
func (spot ParkingSpot) RemoveVehicle() bool {
	spot.free = true
	return true
}

type ParkingSpotAssigner interface {
	Number() string
	IsFree() bool
	Type() int
	SetSpotType(spotType int) bool
	AssignVehicle(vehicle VehicleManger) bool
	RemoveVehicle() bool
}
type HandicappedSpot struct {
	ParkingSpot
}

func (spot ParkingSpot) NewHandicappedSpot() HandicappedSpot {
	s := HandicappedSpot{}

	s.number = spot.number
	s.free = spot.free
	s.vehicle = spot.vehicle
	s.parkingSpotType = spot.parkingSpotType

	return s
}

type CompactSpot struct {
	ParkingSpot
}

func (spot ParkingSpot) NewCompactSpot() CompactSpot {
	s := CompactSpot{}

	s.number = spot.number
	s.free = spot.free
	s.vehicle = spot.vehicle
	s.parkingSpotType = spot.parkingSpotType
	return s
}

type LargeSpot struct {
	ParkingSpot
}

func (spot ParkingSpot) NewLargeSpot() LargeSpot {
	s := LargeSpot{}

	s.number = spot.number
	s.free = spot.free
	s.vehicle = spot.vehicle
	s.parkingSpotType = spot.parkingSpotType
	return s
}

type MotorBikeSpot struct {
	ParkingSpot
}

func (spot ParkingSpot) NewMotorBikeSpot() MotorBikeSpot {
	s := MotorBikeSpot{}

	s.number = spot.number
	s.free = spot.free
	s.vehicle = spot.vehicle
	s.parkingSpotType = spot.parkingSpotType
	return s
}

type ElectricSpot struct {
	ParkingSpot
}

func (spot ParkingSpot) NewElectricSpot() ElectricSpot {
	s := ElectricSpot{}

	s.number = spot.number
	s.free = spot.free
	s.vehicle = spot.vehicle
	s.parkingSpotType = spot.parkingSpotType
	return s
}

/*
Vehicle: Here is the definition for Vehicle and all of its child classes:
*/

type Vehicle struct {
	licensesNumber string
	VehicleType    int
	Ticket         ParkingTicket
}

func (vehicle Vehicle) AssignTicket(ticket ParkingTicket) {
	vehicle.Ticket = ticket
}

type VehicleManger interface {
	AssignTicket(ticket ParkingTicket)
	SetLicensesNumber(licenses string)
	Type() int
}

type Car struct {
	Vehicle
}

func NewCarVehicle() Car {
	v := Car{}
	v.VehicleType = CAR
	return v
}

type Truck struct {
	Vehicle
}

func NewTruckVehicle() Truck {
	v := Truck{}
	v.VehicleType = TRUCK
	return v
}

type Van struct {
	Vehicle
}

func NewVanVehicle() Van {
	v := Van{}
	v.VehicleType = VAN
	return v
}

type Electric struct {
	Vehicle
}

func NewElectricVehicle() Electric {
	v := Electric{}
	v.VehicleType = ELECTRIC
	return v
}

type MotorBike struct {
	Vehicle
}

func NewMotorBikeVehicle() MotorBike {
	v := MotorBike{}
	v.VehicleType = MOTORBIKE
	return v
}

/*
ParkingFloor: This class encapsulates a parking floor:
*/

type ParkingFloor struct {
	name             string
	handicappedSpots map[string]HandicappedSpot
	compactSpots     map[string]CompactSpot
	largeSpots       map[string]LargeSpot
	motorBikeSpots   map[string]MotorBikeSpot
	electricSpots    map[string]ElectricSpot
	infoPortal       map[string]CustomerInfoPortal
	displayBoard     ParkingDisplayBoard
}

func (floor ParkingFloor) SetParkingFloor(name string) {
	floor.name = name
}

func (floor ParkingFloor) AddParkingSpot(spot ParkingSpot) {
	switch spot.Type() {
	case HANDICAPPED_SPOT:
		floor.handicappedSpots[spot.Number()] = spot.NewHandicappedSpot()
		break
	case COMPACT_SPOT:
		floor.compactSpots[spot.Number()] = spot.NewCompactSpot()
		break
	case LARGE_SPOT:
		floor.largeSpots[spot.Number()] = spot.NewLargeSpot()
		break
	case MOTORBIKE_SPOT:
		floor.motorBikeSpots[spot.Number()] = spot.NewMotorBikeSpot()
		break
	case ELECTRIC_SPOT:
		floor.electricSpots[spot.Number()] = spot.NewElectricSpot()
		break
	default:
		log.Printf("not a correct parking spot type %d\n", spot.Type())
	}
}

func (floor ParkingFloor) AssignVehicleToSpot(vehicle VehicleManger, spot ParkingSpotAssigner) {
	spot.AssignVehicle(vehicle)
	switch spot.Type() {
	case HANDICAPPED_SPOT:
	case COMPACT_SPOT:
	case LARGE_SPOT:
	case MOTORBIKE_SPOT:
	case ELECTRIC_SPOT:
		floor.updateDisplayBoardForSpotType(spot)
		break
	default:
		log.Printf("not a correct parking spot type %d\n", spot.Type())
	}
}

func (floor ParkingFloor) updateDisplayBoardForSpotType(spot ParkingSpotAssigner) {
	if _, ok := floor.displayBoard.TakenSpotCountByType[spot.Type()]; ok {
		if spot.IsFree() {
			floor.displayBoard.TakenSpotCountByType[spot.Type()] -= 1
		} else {
			floor.displayBoard.TakenSpotCountByType[spot.Type()] += 1
		}
	}
	floor.displayBoard.showEmptySpotNumbers()
}

func (floor ParkingFloor) freeSpot(spot ParkingSpotAssigner) {
	spot.RemoveVehicle()
	floor.updateDisplayBoardForSpotType(spot)
}

type ParkingDisplayBoard struct {
	TakenSpotCountByType map[int]int
}

func (board ParkingDisplayBoard) showEmptySpotNumbers() {

}

type Ticket struct{}

type ParkingTicket struct{}

type EntrancePanel struct{}
type ExitPanel struct{}

type CustomerInfoPortal struct{}
