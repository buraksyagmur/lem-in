package main

// import (
// 	"fmt"
// 	"strings"
// )

// type room struct {
// 	parent    *room
// 	leftChild *room
// 	rightSib  *room
// 	name      string
// 	occupied  bool
// }

// func Rooms(roomsandConnections []string) {
// 	// var antFarmRooms *room

// 	// fmt.Println(roomsandConnections)
// 	roomNames := []string{}
// 	connections := []string{}
// 	beginConnRmNames := []string{}
// 	endConnRmNames := []string{}

// 	for i := 0; i < len(roomsandConnections); i++ {
// 		for j := 0; j < len(roomsandConnections[i]); j++ {
// 			if roomsandConnections[i][j] == ' ' {
// 				roomName := strings.Split(roomsandConnections[i], " ")[0]
// 				roomNames = append(roomNames, roomName)
// 				break // coz there are 2 spaces
// 			} else if roomsandConnections[i][j] == '-' {
// 				connections = append(connections, roomsandConnections[i])
// 				beginEndSlice := strings.Split(roomsandConnections[i], "-")
// 				beginConnRmNames = append(beginConnRmNames, beginEndSlice[0])
// 				endConnRmNames = append(endConnRmNames, beginEndSlice[1])
// 			}
// 		}
// 	}
// 	// fmt.Println(roomNames)
// 	fmt.Println(connections)
// 	fmt.Println(beginConnRmNames)
// 	fmt.Println(endConnRmNames)
// }
// func addRoom(root *room, rmToAdd room, roomNames, connections []string) *room {
// 	// end room / base case / final case
// 	endRmName := findEndRoomName(roomNames, connections)
// 	if rmToAdd.name == endRmName {
// 		return &room{
// 			// parent: , // set in other cases?
// 			leftChild: nil,
// 			rightSib:  nil,
// 			name:      endRmName,
// 			occupied:  false,
// 		}
// 	}

// 	// from top to bottom?
// 	// for r := 0; r < len(rmToAdd.children); r++ {
// 	// 	root.children[r] = addRoom(root, rmToAdd.children[r], roomNames, connections)
// 	// }

// 	// not start room
// 	// if rmName != roomNames[0] {
// 	// 	for r:=0; r<len(children); r++ {

// 	// 	}
// 	// }
// 	return root

// 	// start room/
// 	// if root == nil {
// 	// 	return &room{
// 	// 		parent: nil,
// 	// 		// children: findChildren(connections),// deal with children in below cases, by refering p
// 	// 		name:     rmName, // subject to change if there is another slice following the order of rooms
// 	// 		occupied: true,
// 	// 	}

// 	// }
// }

// func findEndRoomName(rmNames, connections []string) string {
// 	var endRmName string
// 	for i := 0; i < len(rmNames); i++ {
// 		for c := 0; c < len(connections); c++ {
// 			// if any rmNames[i] is not at any beginning of conn i.e. connections[0]
// 			// then it is the end rm
// 			if rmNames[i] == connections[0] {
// 				continue
// 			} else {
// 				endRmName = rmNames[i]
// 			}
// 		}
// 	}
// 	fmt.Println(endRmName)
// 	return endRmName
// }
