package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/naewatcharapong/hotel-management-test/models"
	"github.com/naewatcharapong/hotel-management-test/util"
)

func main() {
	var myHotel *models.HotelModel
	var outputContent string
	file, err := os.Open("./data/input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "create_hotel":
			floor, _ := strconv.Atoi(parts[1])
			roomsPerFloor, _ := strconv.Atoi(parts[2])
			myHotel = myHotel.NewHotel(floor, roomsPerFloor)
			outputContent += fmt.Sprintf("Hotel created with %d floor(s), %d room(s) per floor.\n", myHotel.Floors, myHotel.RoomsPerFloor)
		case "book":
			room := models.Room{Number: parts[1]}
			name := parts[2]
			age, _ := strconv.Atoi(parts[3])
			currentGuest, room, err := myHotel.CheckIn(room, name, age)
			if err != nil {
				outputContent += fmt.Sprintf("Cannot book room %s for %s, The room is currently booked by %s.\n", room.Number, name, currentGuest)
			} else {
				outputContent += fmt.Sprintf("Room %s is booked by %s with keycard number %d.\n", room.Number, name, room.Guest.Keycard)
			}
		case "list_available_rooms":
			availableRooms := myHotel.AvailableRooms()
			for _, room := range availableRooms {
				outputContent += fmt.Sprintf("%s\n",room.Number)
			}
		case "checkout":
			keycard, _ := strconv.Atoi(parts[1])
			room, wrongGuest, err := myHotel.CheckOut(keycard, parts[2])
			if err != nil {
				outputContent += fmt.Sprintf("Only %s can checkout with keycard number %d.\n", wrongGuest.Name, wrongGuest.Keycard)
			} else {
				outputContent += fmt.Sprintf("Room %s is checkout\n", room)
			}
		case "list_guest":
			guests := myHotel.GuestList()
			var listGuest []string
			for _, guest := range guests {
				listGuest = append(listGuest, guest.Name)
			}
			outputContent += fmt.Sprintf("%s\n", strings.Join(listGuest, ", "))
		case "get_guest_in_room":
			guestInRoom := myHotel.GuestInRoom(parts[1])
			outputContent += fmt.Sprintf("%s\n", guestInRoom.Name)
		case "list_guest_by_age":
			var listGuestByAge []string
			ageNumber, _ := strconv.Atoi(parts[2])
			guestsInAgeRange := myHotel.GuestListByAge(parts[1], ageNumber)

			for _, guest := range guestsInAgeRange {
				listGuestByAge = append(listGuestByAge, guest.Name)
			}
			outputContent += fmt.Sprintf("%s\n", strings.Join(listGuestByAge, ", "))
		case "list_guest_by_floor":
			var listGuestByFloor []string
			guestsInFloorRange := myHotel.GuestListByFloor(parts[1])
			for _, guest := range guestsInFloorRange {
				listGuestByFloor = append(listGuestByFloor, guest.Name)
			}
			outputContent += fmt.Sprintf("%s\n", strings.Join(listGuestByFloor, ", "))
		case "checkout_guest_by_floor":
			listRoomInFloorRange, _ := myHotel.CheckOutByFloor(parts[1])

			outputContent += fmt.Sprintf("Room %s are checkout.\n", strings.Join(listRoomInFloorRange, ", "))
		case "book_by_floor":
			var listRoomByFloor []string
			var listKeyCardByFloor []int
			age, _ := strconv.Atoi(parts[3])
			listCheckinRoomInFloor, err := myHotel.CheckInByFloor(parts[1], parts[2], age)
			if err != nil {
				outputContent += fmt.Sprintf("Cannot book floor %s for %s.\n", parts[1], parts[2])
				break
			}
			for _, room := range listCheckinRoomInFloor {
				listRoomByFloor = append(listRoomByFloor, room.Number)
				listKeyCardByFloor = append(listKeyCardByFloor, room.Guest.Keycard)
			}

			outputContent += fmt.Sprintf("Room %s are booked with keycard number %s\n", strings.Join(listRoomByFloor, ", "), (util.ArrayToString(listKeyCardByFloor, ", ")))
		}
	}
	util.WriteOutputFile(outputContent)
}
