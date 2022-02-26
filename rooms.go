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

func Rooms(roomsandConnections []string) {
	antFarmRooms := []room{}

	// fmt.Println(slccontent)
	// fmt.Println(roomsandConnections)
	// fmt.Println(RoomsandConnections[0])
	// fmt.Println(RoomsandConnections[len(RoomsandConnections)-1])
	roomNames := []string{}
	connections := []string{}
	// start room name
	// roomNames = append(roomNames, strings.Split(slccontent[startline+1], " ")[0])
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
	fmt.Println(roomNames)
	fmt.Println(nil)
	fmt.Println(connections)

	startingRoom := room{
		name:   roomNames[0],
		parent: nil,
		// children: startChildrenRm,
		occupied: true,
	}

	startChildrenName := []string{}
	for c := 0; c < len(connections); c++ {
		beginRm := strings.Split(connections[c], "-")
		if beginRm[0] == startingRoom.name {
			startChildrenName = append(startChildrenName, beginRm[0])
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

func createOtherRoom(name string, p *room, children []room) {
	newRoom := room{
		name:     name,
		parent:   p,
		children: children,
		occupied: false,
	}
	fmt.Print(newRoom)
}

func findEndRoom(rmNames, connections []string) {
	for i := 0; i < len(rmNames); i++ {
		for c := 0; c < len(connections); c++ {
			// if any rmNames[i] is not at any beginning of conn i.e. connections[0]
			// then it is the end rm
		}
	}
}
