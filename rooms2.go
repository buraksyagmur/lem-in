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
	endConnRmNames := []string{}

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
				endConnRmNames = append(endConnRmNames, beginEndSlice[1])
			}
		}
	}
	// fmt.Println(roomNames)
	fmt.Println(connections)
	fmt.Println(beginConnRmNames)
	fmt.Println(endConnRmNames)

	endRmName := findEndRoomName(roomNames, connections)

	// test cases
	antFarmRooms = addRoom(antFarmRooms, beginConnRmNames[0], endConnRmNames[0], true, roomNames[0], endRmName)
	antFarmRooms = addRoom(antFarmRooms, endConnRmNames[0], endConnRmNames[0], true, roomNames[0], endRmName)

	// for loop
	antFarmRooms = addRoom(antFarmRooms, beginConnRmNames[0], endConnRmNames[0], true, roomNames[0], endRmName)
}
func addRoom(root *room, rmToAddName, childRmName string, occupied bool, startRmName, endRmName string) *room {
	// end room / base case / final case
	if rmToAddName == endRmName {
		return &room{
			// parent: ,
			children: nil,
			name:     endRmName,
			occupied: false,
		}
	} else if rmToAddName == startRmName { // start Room special case
		roomToAdd := room{
			parent: nil,
			// children: startChildrenRm,
			name:     startRmName,
			occupied: true,
		}
	} else {
		roomToAdd := room{
			// parent: nil,
			// children: startChildrenRm,
			name:     rmToAddName,
			occupied: false,
		}
	}

	childrenName := []string{}
	for c := 0; c < len(connections); c++ {
		connectionsRms := strings.Split(connections[c], "-")
		beginRm := connectionsRms[0]
		// destinationRm := connectionsRms[1] // have to match the right begin Room name
		if beginRm == startingRoom.name {
			startChildrenName = append(startChildrenName, beginRm)
			// startGrandChildrenName = append(startGrandChildrenName, destinationRm)
		}
	}

	// 	startChildrenRm := []*room{}
	// 	for s := 0; s < len(startChildrenName); s++ {
	// 		startChildrenRm = append(startChildrenRm, &room{
	// 			parent: &startingRoom,
	// 			// children: ,
	// 			name:     startChildrenName[s],
	// 			occupied: false,
	// 		})
	// 	}

	// 	startingRoom.children = startChildrenRm

	// from top to bottom?
	// childrenSlice := make([]room, 10)
	// for r := 0; r < len(rmToAdd.children); r++ {
	// 	root.children[r] = addRoom(root, rmToAdd.children[r], roomNames, connections)
	// }

	// not start room
	// if rmName != roomNames[0] {
	// 	for r:=0; r<len(children); r++ {

	// 	}
	// }
	return root

	// start room/
	// if root == nil {
	// 	return &room{
	// 		parent: nil,
	// 		// children: findChildren(connections),// deal with children in below cases, by refering p
	// 		name:     rmName, // subject to change if there is another slice following the order of rooms
	// 		occupied: true,
	// 	}

	// }
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
	fmt.Println(endRmName)
	return endRmName
}

func printRoom(root *room) {
	if root == nil {
		fmt.Println("No rooms")
		return
	}
	fmt.Println(root.name)
	i := 0
	for root.children != nil {
		printRoom(root.children[i])
		i++
	}

}
