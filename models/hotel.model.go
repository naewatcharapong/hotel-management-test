package models

import (
	"fmt"
	"strings"
)

type HotelModel struct {
	Floors        int
	RoomsPerFloor int
	Rooms         []Room
}

// Create new hotel.
func (h *HotelModel) NewHotel(floors, roomsPerFloor int) *HotelModel {
	rooms := make([]Room, floors*roomsPerFloor)
	for i := 0; i < floors; i++ {
		for j := 0; j < roomsPerFloor; j++ {
			roomNumber := fmt.Sprintf("%d%02d", i+1, j+1)
			rooms[i*roomsPerFloor+j] = Room{Number: roomNumber}
		}
	}
	return &HotelModel{Floors: floors, RoomsPerFloor: roomsPerFloor, Rooms: rooms}
}

// Generate key card for each room.
func (h *HotelModel) GenerateKeyCard() int {
	var intitailKeyCard int = 1
	for i := 0; i < len(h.Rooms); i++ {
		for _, room := range h.Rooms {
			if room.Guest.Keycard == intitailKeyCard {
				intitailKeyCard++
				break
			}
		}
	}
	return intitailKeyCard

}

// Check-in function
func (h *HotelModel) CheckIn(r Room, guestName string, guestAge int) (string, Room, error) {
	var currentGuest string = ""
	for i, room := range h.Rooms {
		if room.Number == r.Number && room.Guest.Name == "" {
			h.Rooms[i].Guest = GuestModel{Name: guestName, Age: guestAge, Keycard: h.GenerateKeyCard()}
			return "", h.Rooms[i], nil
		}
		if room.Number == r.Number && room.Guest.Name != guestName {
			currentGuest = room.Guest.Name
			return currentGuest, Room{}, fmt.Errorf("room not found or not avilable")
		}

	}

	return "", Room{}, fmt.Errorf("room not found or not available")
}

// Check-in by floor function
func (h *HotelModel) CheckInByFloor(floor, name string, age int) ([]Room, error) {
	var listCheckInRoom []Room
	for _, room := range h.Rooms {
		if room.Guest.Name != "" && strings.HasPrefix(room.Number, floor) && room.Guest.Keycard != 0 {
			return listCheckInRoom, fmt.Errorf("floor is not available")
		}
		if room.Guest.Name == "" && strings.HasPrefix(room.Number, floor) && room.Guest.Keycard == 0 {
			_, room, err := h.CheckIn(room, name, age)
			if err != nil {
				return listCheckInRoom, fmt.Errorf("room not found or not available")
			}
			listCheckInRoom = append(listCheckInRoom, room)
		}
	}
	return listCheckInRoom, nil
}

// Check-out function
func (h *HotelModel) CheckOut(keycard int, name string) (string, GuestModel, error) {
	var wrongGuest GuestModel = GuestModel{}
	for i, room := range h.Rooms {
		if room.Guest.Keycard == keycard && room.Guest.Name == name {
			h.Rooms[i].Guest = GuestModel{}
			return h.Rooms[i].Number, wrongGuest, nil
		}
		if room.Guest.Keycard == keycard && room.Guest.Name != name {
			wrongGuest.Name = room.Guest.Name
			wrongGuest.Keycard = room.Guest.Keycard
			return h.Rooms[i].Number, wrongGuest, fmt.Errorf("keycard not found")
		}
	}
	return "", wrongGuest, fmt.Errorf("keycard not found")
}

// Check-out by floor function
func (h *HotelModel) CheckOutByFloor(floor string) ([]string, error) {
	var listCheckOutRoom []string
	for _, room := range h.Rooms {
		if room.Guest.Name != "" && strings.HasPrefix(room.Number, floor) && room.Guest.Keycard != 0 {
			h.CheckOut(room.Guest.Keycard, room.Guest.Name)
			listCheckOutRoom = append(listCheckOutRoom, room.Number)
		}
	}
	return listCheckOutRoom, nil
}

// Get list of available rooms.
func (h *HotelModel) AvailableRooms() []Room {
	var availableRooms []Room
	for _, room := range h.Rooms {
		if room.Guest.Name == "" {
			availableRooms = append(availableRooms, room)
		}
	}
	return availableRooms
}

// Get list of user.
func (h *HotelModel) GuestList() []GuestModel {
	var guests []GuestModel
	for _, room := range h.Rooms {
		if room.Guest.Name != "" {
			guests = append(guests, room.Guest)
		}
	}
	return guests
}

// Get list of user by age.
func (h *HotelModel) GuestListByAge(Value string, Age int) []GuestModel {
	var guests []GuestModel
	for _, room := range h.Rooms {
		if room.Guest.Name != "" && Value == "<" && room.Guest.Age < Age {
			guests = append(guests, room.Guest)
		}
		if room.Guest.Name != "" && Value == ">" && room.Guest.Age > Age {
			guests = append(guests, room.Guest)
		}

	}
	return guests
}

// Get list of user by floor.
func (h *HotelModel) GuestListByFloor(floor string) []GuestModel {
	var guests []GuestModel
	for _, room := range h.Rooms {

		if room.Guest.Name != "" && strings.HasPrefix(room.Number, floor) && room.Guest.Keycard != 0 {
			guests = append(guests, room.Guest)
		}
	}
	return guests
}

// Get the user in room.
func (h *HotelModel) GuestInRoom(roomNumber string) GuestModel {
	for _, room := range h.Rooms {
		if room.Number == roomNumber {
			return room.Guest
		}
	}
	return GuestModel{}
}
