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
	antFarmRooms := []room{}

	// fmt.Println(roomsandConnections)
	roomNames := []string{}
	connections := []string{}

	for i := 0; i < len(roomsandConnections); i++ {
		for j := 0; j < len(roomsandConnections[i]); j++ {
			if roomsandConnections[i][j] == ' ' {
				roomName := strings.Split(roomsandConnections[i], " ")[0]
				roomNames = append(roomNames, roomName)
				break // coz there are 2 spaces
			} else if roomsandConnections[i][j] == '-' {
				connections = append(connections, roomsandConnections[i])
			}
		}
	}
	// fmt.Println(roomNames)
	// fmt.Println(connections)

	// start Room special case
	startingRoom := room{
		name:   roomNames[0],
		parent: nil,
		// children: startChildrenRm,
		occupied: true,
	}

	startChildrenName := []string{}
	// startGrandChildrenName := []string{}
	for c := 0; c < len(connections); c++ {
		connectionsRms := strings.Split(connections[c], "-")
		beginRm := connectionsRms[0]
		// destinationRm := connectionsRms[1] // have to match the right begin Room name
		if beginRm == startingRoom.name {
			startChildrenName = append(startChildrenName, beginRm)
			// startGrandChildrenName = append(startGrandChildrenName, destinationRm)
		}
	}

	startChildrenRm := []*room{}
	for s := 0; s < len(startChildrenName); s++ {
		startChildrenRm = append(startChildrenRm, &room{
			parent: &startingRoom,
			// children: ,
			name:     startChildrenName[s],
			occupied: false,
		})
	}

	startingRoom.children = startChildrenRm

	antFarmRooms = append(antFarmRooms, startingRoom)
	// antFarmRooms = append(antFarmRooms, startChildrenRooms)

	// test addRoom func endRoom case
	// antFarmRooms = append(antFarmRooms, *addRoom(&room{
	// 	name: "peter",
	// }, roomNames, connections))

	// antFarmRooms = append(antFarmRooms, *addRoom(&room{
	// 	name: "peter",
	// }, roomNames, connections))

	findEndRoomName(roomNames, connections)
	fmt.Println("ant Room: ", antFarmRooms)
}

// func createOtherRoom(name string, p *room, children []room) {
// 	newRoom := room{
// 		name:     name,
// 		parent:   p,
// 		children: children,
// 		occupied: false,
// 	}
// 	fmt.Print(newRoom)
// }

// func findChildren(connections []string, cur *room) []*room {
// 	childrenName := []string{}
// 	for c := 0; c < len(connections); c++ {
// 		beginRm := strings.Split(connections[c], "-")[0]
// 		if beginRm == (*cur).name {
// 			childrenName = append(childrenName, beginRm)
// 		}

// 	}
// 	fmt.Print(childrenName)

// 	childrenRm := []*room{}
// 	for s := 0; s < len(childrenName); s++ {
// 		childrenRm = append(childrenRm, &room{
// 			name:   childrenName[s],
// 			parent: cur,
// 			// children: ,
// 			occupied: false,
// 		})
// 	}
// 	fmt.Print(childrenRm)
// 	(*cur).children = childrenRm
// }

// roomNames should be replaced by a slice that is ordered by connections
func addRoom(root *room, rmName string, roomNames, connections []string) *room {
	// end room / base case / final case
	endRmName := findEndRoomName(roomNames, connections)
	if rmName == endRmName {
		return &room{
			// parent: , // set in other cases?
			children: nil,
			name:     endRmName,
			occupied: false,
		}
	}
	// not start room
	// if rmName != roomNames[0] {
	// 	for r:=0; r<len(?); r++ {

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
