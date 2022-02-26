package main

import (
	"fmt"
	"strings"
)

type room struct {
	name     string
	parent   *room
	children []room
	occupied bool
}

func Rooms() {
	antFarmRooms := []room{}

	startingRoom := room{
		name:   strings.Split(slccontent[startline+1], " ")[0],
		parent: nil,
		// children: startChildrenRm,
		occupied: true,
	}

	startChildrenName := []string{}
	for m := connectionStartLine; m < len(RoomsandConnections); m++ {
		connections := strings.Split(RoomsandConnections[m], "-")
		if connections[0] == startingRoom.name {
			startChildrenName = append(startChildrenName, connections[1])
		}
	}
	// fmt.Print(startChildrenName)

	startChildrenRm := []room{}
	for s := 0; s < len(startChildrenName); s++ {
		startChildrenRm = append(startChildrenRm, room{
			name:   startChildrenName[s],
			parent: &startingRoom,
			// children: ,
			occupied: false,
		})
	}
	// fmt.Print(startChildrenRm)
	startingRoom.children = startChildrenRm

	antFarmRooms = append(antFarmRooms, startingRoom)
	// make rooms with createRoom
	fmt.Println(antFarmRooms)
}
