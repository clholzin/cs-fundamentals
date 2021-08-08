package parkinglot

import (
	"fmt"
	"log"
	"strings"
)

/*
Following is the skeleton code for our parking lot system:

Enums and Constants: Here are the required enums, data types, and constants:
*/
// TotalFloors
const (
	TotalFloors = 3
)

// VehicleType
const (
	CAR = iota
	TRUCK
	ELECTRIC
	VAN
	MOTORBIKE
)

// ParkingSpotType
const (
	TypeHandicappedSpot = iota
	TypeCompactSpot
	TypeLargeSpot
	TypeMotorBikeSpot
	TypeElectricSpot
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
	ActiveTicket = iota
	PaidTicket
	LostTicket
)

// FloorTotals
const (
	FloorTotalHandicapped = 5
	FloorTotalCompact     = 12
	FloorTotalLarge       = 10
	FloorTotalMotorBike   = 2
	FloorTotalElectric    = 5
)

// Address ...
type Address struct {
	StreetAddress string
	City          string
	State         string
	ZipCode       string
	Country       string
}

// Person ...
type Person struct {
	Name    string
	Address Address
	Email   string
	Phone   string
}

/*
Account, Admin, and ParkingAttendant: These classes represent various people that interact with our system:
*/

// Account ...
type Account struct {
	userName string
	password string
	status   string
	Person
}

// Accounter ...
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
	case TypeHandicappedSpot:
	case TypeCompactSpot:
	case TypeLargeSpot:
	case TypeMotorBikeSpot:
	case TypeElectricSpot:
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

// NewVehicle ...
func NewVehicle(cartype int) (VehicleManger, error) {
	switch cartype {
	case CAR:
		return NewCarVehicle(), nil
	case TRUCK:
		return NewTruckVehicle(), nil
	case VAN:
		return NewVanVehicle(), nil
	case MOTORBIKE:
		return NewMotorBikeVehicle(), nil
	case ELECTRIC:
		return NewElectricVehicle(), nil
	}
	return nil, fmt.Errorf("wrong type provided")
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

// ParkingFloor: This class encapsulates a parking floor:
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
	case TypeHandicappedSpot:
		floor.handicappedSpots[spot.Number()] = spot.NewHandicappedSpot()
		break
	case TypeCompactSpot:
		floor.compactSpots[spot.Number()] = spot.NewCompactSpot()
		break
	case TypeLargeSpot:
		floor.largeSpots[spot.Number()] = spot.NewLargeSpot()
		break
	case TypeMotorBikeSpot:
		floor.motorBikeSpots[spot.Number()] = spot.NewMotorBikeSpot()
		break
	case TypeElectricSpot:
		floor.electricSpots[spot.Number()] = spot.NewElectricSpot()
		break
	default:
		log.Printf("not a correct parking spot type %d\n", spot.Type())
	}
}

func (floor ParkingFloor) AssignVehicleToSpot(vehicle VehicleManger, spot ParkingSpotAssigner) {
	spot.AssignVehicle(vehicle)
	switch spot.Type() {
	case TypeHandicappedSpot:
	case TypeCompactSpot:
	case TypeLargeSpot:
	case TypeMotorBikeSpot:
	case TypeElectricSpot:
		floor.updateDisplayBoardForSpotType(spot)
		break
	default:
		log.Printf("not a correct parking spot type %d\n", spot.Type())
	}
}

func (floor ParkingFloor) updateDisplayBoardForSpotType(spot ParkingSpotAssigner) {

	if spot.IsFree() {
		floor.displayBoard.TakenSpotCountByType[spot.Type()] -= 1
	} else {
		floor.displayBoard.TakenSpotCountByType[spot.Type()] += 1
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

func NewParkingDisplayBoard() ParkingDisplayBoard {
	display := ParkingDisplayBoard{}
	display.TakenSpotCountByType[TypeHandicappedSpot] = 0
	display.TakenSpotCountByType[TypeCompactSpot] = 0
	display.TakenSpotCountByType[TypeLargeSpot] = 0
	display.TakenSpotCountByType[TypeMotorBikeSpot] = 0
	display.TakenSpotCountByType[TypeElectricSpot] = 0

	return display
}

type FloorFreeSpotCounts struct {
	Handicapped int
	Compact     int
	Large       int
	Electric    int
	MotorBike   int
}

func (board ParkingDisplayBoard) GetEmptySpotNumbers() FloorFreeSpotCounts {

	return FloorFreeSpotCounts{
		Handicapped: FloorTotalHandicapped - board.TakenSpotCountByType[TypeHandicappedSpot],
		Compact:     FloorTotalCompact - board.TakenSpotCountByType[TypeCompactSpot],
		Large:       FloorTotalLarge - board.TakenSpotCountByType[TypeLargeSpot],
		MotorBike:   FloorTotalMotorBike - board.TakenSpotCountByType[TypeMotorBikeSpot],
		Electric:    FloorTotalElectric - board.TakenSpotCountByType[TypeElectricSpot],
	}
}

func (board ParkingDisplayBoard) showEmptySpotNumbers() {

	display := strings.Builder{}

	freeHandicapped := FloorTotalHandicapped - board.TakenSpotCountByType[TypeHandicappedSpot]

	if freeHandicapped > 0 {
		display.WriteString(fmt.Sprintf("Free Handicapped: %d", freeHandicapped))
	} else {
		display.WriteString("Handicapped Full")
	}

	freeCompact := FloorTotalCompact - board.TakenSpotCountByType[TypeCompactSpot]

	if freeCompact > 0 {
		display.WriteString(fmt.Sprintf("Free Compact: %d", freeCompact))
	} else {
		display.WriteString("Compact Full")
	}

	freeLarge := FloorTotalLarge - board.TakenSpotCountByType[TypeLargeSpot]

	if freeLarge > 0 {
		display.WriteString(fmt.Sprintf("Free Large: %d", freeLarge))
	} else {
		display.WriteString("Large Full")
	}

	freeMotorBike := FloorTotalMotorBike - board.TakenSpotCountByType[TypeMotorBikeSpot]

	if freeMotorBike > 0 {
		display.WriteString(fmt.Sprintf("Free MotorBike: %d", freeMotorBike))
	} else {
		display.WriteString("MotorBike Full")
	}

	freeElectric := FloorTotalElectric - board.TakenSpotCountByType[TypeElectricSpot]

	if freeElectric > 0 {
		display.WriteString(fmt.Sprintf("Free Electric: %d", freeElectric))
	} else {
		display.WriteString("Electric Full")
	}

	log.Println(display.String())
}

/*
ParkingLot: Our system will have only one object of this class. This can be enforced by using the Singleton pattern. In software engineering, the singleton pattern is a software design pattern that restricts the instantiation of a class to only one object.
*/

type ParkingLot struct {
	Name         string
	Address      Location
	Rate         ParkingRate
	MaxCompact   int
	MaxLarge     int
	MaxMotorBike int
	MaxElectric  int

	// refactor current and max counts

	EntrancePanels map[string]EntrancePanel
	ExitPanels     map[string]ExitPanel
	ParkingFloors  map[string]ParkingFloor

	ActiveTickets map[string]ParkingTicket
}

func NewParkingLot() ParkingLot {
	return ParkingLot{
		MaxHandicapped: TotalFloors * FloorTotalHandicapped,
		MaxCompact:     TotalFloors * FloorTotalCompact,
		MaxLarge:       TotalFloors * FloorTotalLarge,
		MaxMotorBike:   TotalFloors * FloorTotalMotorBike,
		MaxElectric:    TotalFloors * FloorTotalElectric,
	}
}

func (lot ParkingLot) NewParkingTicket(vehicle VehicleManger) (Ticket, error) {
	if lot.IsFullBy(vehicle.Type()) {
		return nil, fmt.Errorf("Lot Full for type")
	}
	ticket := Ticket{}
	vehicle.AssignTicket(ticket)
	ticket.Save()

	lot.IncrementSpotCount(vehicle.Type)
	lot.ActiveTickets[ticket.Number()] = ticket
	return ticket, nil
}

func (lot ParkingLot) IsFullBy(vehicleType int) bool {
	totals := EmptySpotsTotals()
	switch vehicleType {
	case Car:
		return (totals.Compact + totals.Large) >= (lot.MaxCompact + lot.MaxLarge)
	case TRUCK:
	case VAN:
		return totals.Large >= lot.MaxLarge
	case MOTORBIKE:
		return totals.MotorBike >= lot.MaxMotorBike
	case ELECTRIC:
		return (totals.Compact + totals.Electric) >= (lot.Compact + lot.Electric)
	default:
		log.Printf("not a correct type %d\n", vehicleType)
	}
	return false
}

func (lot ParkingLot) EmptySpotsTotals() FloorFreeSpotCounts {

	totals := FloorFreeSpotCounts{}
	for _, floor := range lot.ParkingFloors {
		floorCounts := floor.GetEmptySpotNumbers()
		totals.Handicapped += floorCounts.Handicapped
		totals.Compact += floorCounts.Compact
		totals.Large += floorCounts.Large
		totals.MotorBike += floorCounts.MotorBike
		totals.Electric += floorCounts.Electric
	}
	return totals
}
func (lot ParkingLot) IncrementSpotCount(vehicleType int) {}

type ParkingRate struct{}

type Location struct{}

type Ticket struct {
	Number string
}

func (ticket Ticket) Save() {}

type ParkingTicket struct{}

type EntrancePanel struct{}
type ExitPanel struct{}

type CustomerInfoPortal struct{}
