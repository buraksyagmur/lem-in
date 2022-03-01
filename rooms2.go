package main

import (
	"fmt"
	"strings"
)

type room2 struct {
	parent   *room
	children [](*room)
	name     string
	occupied bool
}

func Rooms2(roomsandConnections []string) {
	// var antFarmRooms *room

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

}
func addRoom2(root *room, rmName string, roomNames, connections []string) *room {
	// end room / base case / final case
	endRmName := findEndRoomName2(roomNames, connections)
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

func findEndRoomName2(rmNames, connections []string) string {
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
