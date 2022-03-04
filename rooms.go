package main

import (
	"fmt"
	"strings"
)

type room struct {
	parent   *room
	children [](*room)
	name     string
	occupied bool
}

func Rooms(roomsandConnections []string) {
	var antFarmRooms *room

	// fmt.Println(roomsandConnections)
	roomNames := []string{}
	connections := []string{}
	beginConnRmNames := []string{}
	destConnRmNames := []string{}

	for i := 0; i < len(roomsandConnections); i++ {
		for j := 0; j < len(roomsandConnections[i]); j++ {
			if roomsandConnections[i][j] == ' ' {
				roomName := strings.Split(roomsandConnections[i], " ")[0]
				roomNames = append(roomNames, roomName)
				break // coz there are 2 spaces
			} else if roomsandConnections[i][j] == '-' {
				connections = append(connections, roomsandConnections[i])
				beginEndSlice := strings.Split(roomsandConnections[i], "-")
				beginConnRmNames = append(beginConnRmNames, beginEndSlice[0])
				destConnRmNames = append(destConnRmNames, beginEndSlice[1])
			}
		}
	}
	// fmt.Println(roomNames)
	// fmt.Println(connections)
	// fmt.Println(beginConnRmNames)
	// fmt.Println(destConnRmNames)

	startRmName := roomNames[0]
	endRmName := findEndRoomName(roomNames, connections)

	antFarmRooms = addRoom(antFarmRooms, startRmName, startRmName, endRmName, beginConnRmNames, destConnRmNames)

	// test cases
	// first name in roomNames must be the start room
	// for r := 0; r < len(roomNames); r++ {
	// 	antFarmRooms = addRoom(antFarmRooms, roomNames[r], startRmName, endRmName, beginConnRmNames, destConnRmNames)
	// }

	// antFarmRooms = addRoom(antFarmRooms, roomNames[len(connections)-1], true, startRmName, endRmName)

	// loop?
	// antFarmRooms = addRoom(antFarmRooms, childrenName, endConnRmNames[0], true, roomNames[0], endRmName)

	fmt.Println("----------------")
	fmt.Println(antFarmRooms)
	fmt.Println("----------------")
	printRoom(antFarmRooms)
}
func addRoom(root *room, rmToAddName string, startRmName, endRmName string, beginConnRmNames, destConnRmNames []string) *room {
	var roomToAdd *room
	// end room / base case / final case
	if rmToAddName == endRmName {
		fmt.Println("___________________________________________________")
		fmt.Printf("constructing end room %s...\n", rmToAddName)
		return &room{
			// parent: ,
			children: nil,
			name:     rmToAddName,
			occupied: false,
		}
	} else if rmToAddName == startRmName { // start Room special case
		fmt.Println("___________________________________________________")
		fmt.Printf("constructing start room %s...\n", rmToAddName)
		roomToAdd = &room{
			parent: nil,
			// children: startChildrenRm,
			name:     rmToAddName,
			occupied: true,
		}
	} else {
		fmt.Println("___________________________________________________")
		fmt.Printf("constructing other room %s...\n", rmToAddName)
		roomToAdd = &room{
			// parent: nil,
			// children: startChildrenRm,
			name:     rmToAddName,
			occupied: false,
		}
	}

	childrenRm := []*room{}
	for c := 0; c < len(beginConnRmNames); c++ {
		beginRmName := beginConnRmNames[c]
		if beginRmName == rmToAddName {
			destRmName := destConnRmNames[c]
			fmt.Printf("Adding the %dth child (%s) to %s\n", c, destRmName, rmToAddName)
			childrenRm = append(childrenRm, &room{
				parent: roomToAdd,
				// children: ,
				name:     destRmName,
				occupied: false,
			})
			addRoom(roomToAdd, destRmName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
		}
	}
	return root
}

func findEndRoomName(rmNames, connections []string) string {
	var endRmName string
	for i := 0; i < len(rmNames); i++ {
		for c := 0; c < len(connections); c++ {
			// if any rmNames[i] is not at any beginning of conn i.e. connections[0]
			// then it is the end rm
			if rmNames[i] == connections[0] {
				continue
			} else {
				endRmName = rmNames[i]
			}
		}
	}
	// fmt.Println(endRmName)
	return endRmName
}

func printRoom(root *room) {
	fmt.Println("-------Ant--Farm--Rooms--------")
	if root == nil {
		fmt.Println("No rooms")
		return
	}
	// fmt.Println(root.name)
	i := 0
	for root.children != nil {
		printRoom(root.children[i])
		i++
	}
}
