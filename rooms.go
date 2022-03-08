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

var (
	firstRm *room
	lastRm  *room
)

func Rooms(roomsandConnections []string) *room {
	var antFarmRooms *room

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
			}
		}
	}
	startRmName := roomNames[0]
	endRmName := roomNames[len(roomNames)-1]

	for i := 0; i < len(roomsandConnections); i++ {
		for j := 0; j < len(roomsandConnections[i]); j++ {
			if roomsandConnections[i][j] == '-' {
				connections = append(connections, roomsandConnections[i])
				beginDestSlice := strings.Split(roomsandConnections[i], "-")
				if beginDestSlice[0] == endRmName {
					beginConnRmNames = append(beginConnRmNames, beginDestSlice[1])
					destConnRmNames = append(destConnRmNames, beginDestSlice[0])
				}
				beginConnRmNames = append(beginConnRmNames, beginDestSlice[0])
				destConnRmNames = append(destConnRmNames, beginDestSlice[1])
			}
		}
	}
	antFarmRooms = addRoom(antFarmRooms, startRmName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
	return antFarmRooms
}

func findChildren(roomToAdd *room, rmToAddName string, startRmName, endRmName string, beginConnRmNames, destConnRmNames []string) {
	childrenRm := []*room{}
	for c := 0; c < len(beginConnRmNames); c++ {
		beginRmName := beginConnRmNames[c]
		if beginRmName == rmToAddName {
			destRmName := destConnRmNames[c]
			fmt.Printf("Adding new child room (%s) from the %dth connection to %s\n", destRmName, c, rmToAddName)
			childrenRm = append(childrenRm, &room{
				parent:   roomToAdd,
				name:     destRmName,
				occupied: false,
			})
			addRoom(roomToAdd, destRmName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
		}
	}
}

func addRoom(root *room, rmToAddName string, startRmName, endRmName string, beginConnRmNames, destConnRmNames []string) *room {
	var roomToAdd *room
	if rmToAddName == endRmName { // end room / base case
		fmt.Println("___________________________________________________")
		fmt.Printf("constructing end room %s...\n", rmToAddName)
		fmt.Println("")
		fmt.Println("")
		lastRm = &room{
			parent:   root,
			children: nil,
			name:     rmToAddName,
			occupied: false,
		}
		return lastRm
	} else if rmToAddName == startRmName { // start Room special case
		fmt.Println("___________________________________________________")
		fmt.Printf("constructing start room %s...\n", rmToAddName)
		roomToAdd = &room{
			parent:   nil,
			name:     rmToAddName,
			occupied: true,
		}
		findChildren(roomToAdd, rmToAddName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
		firstRm = roomToAdd
	} else {
		fmt.Println("___________________________________________________")
		fmt.Printf("constructing other room %s...\n", rmToAddName)
		roomToAdd = &room{
			parent:   root,
			name:     rmToAddName,
			occupied: false,
		}
		roomToAdd.parent.children = append(roomToAdd.parent.children, roomToAdd)
		findChildren(roomToAdd, rmToAddName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
	}
	// fmt.Println("lastrm", lastRm)
	// fmt.Println("firstrm", firstRm)
	// fmt.Println(roomToAdd)
	return roomToAdd
}

// func findEndRoomName(rmNames, connections []string) string {
//     var endRmName string
//     for i := 0; i < len(rmNames); i++ {
//         for c := 0; c < len(connections); c++ {
//             // if any rmNames[i] is not at any beginning of conn i.e. connections[0]
//             // then it is the end rm
//             if rmNames[i] == connections[0] {
//                 continue
//             } else {
//                 endRmName = rmNames[i]
//             }
//         }
//     }
//     // fmt.Println(endRmName)
//     return endRmName
// }
// func printRoom(root *room) { // not working yet, but not required
// 	fmt.Println("-------Ant--Farm--Rooms--------")
// 	if root == nil {
// 		return
// 	}
// 	fmt.Println(root.name)
// 	for i := 0; i < len(root.children); i++ {
// 		printRoom(root.children[i])
// 	}
// }
