package main

import (
	"fmt"
	"strings"
)

type room struct {
	parent   [](*room)
	name     string
	children [](*room)
	occupied bool
}

var (
	firstRm  *room
	lastRm   *room
	Farm     []room
	test     []*room
	count    int
	checkdup []string
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
	startRmName := nameofStart
	endRmName := nameofEnd

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
	// fmt.Println("roomname", destConnRmNames)
	antFarmRooms = addRoom(antFarmRooms, startRmName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
	return antFarmRooms
}

func findChildren(roomToAdd *room, rmToAddName string, startRmName, endRmName string, beginConnRmNames, destConnRmNames []string) {
	// childrenRm := []*room{}
	for c := 0; c < len(beginConnRmNames); c++ {
		beginRmName := beginConnRmNames[c]
		if beginRmName == rmToAddName {
			// fmt.Println("dsf", childrenRm)
			destRmName := destConnRmNames[c]
			fmt.Printf("Adding new child room (%s) from the %dth connection to %s\n", destRmName, c, rmToAddName)
			// fmt.Println(childrenRm)
			// childrenRm = append(childrenRm, &room{
			// 	parent:   roomToAdd,
			// 	name:     destRmName,
			// 	occupied: false,
			// })
			// fmt.Println("2nd", childrenRm)
			// fmt.Println("children", roomToAdd, destRmName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
			// fmt.Println(childrenRm)
			fmt.Println("infos", roomToAdd, destRmName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
			addRoom(roomToAdd, destRmName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
		}
	}
}

func addRoom(root *room, rmToAddName string, startRmName, endRmName string, beginConnRmNames, destConnRmNames []string) *room {
	var roomToAdd *room
	checkdup = append(checkdup, rmToAddName)
	if rmToAddName == endRmName { // end room / base case
		fmt.Println("___________________________________________________")
		fmt.Printf("constructing end room %s...\n", rmToAddName)
		test = nil
		test = append(test, root)
		fmt.Println(rmToAddName)
		if dup(checkdup) != rmToAddName {
			lastRm = &room{
				parent:   test,
				name:     rmToAddName,
				children: nil,
				occupied: false,
			}
		} else {
			fmt.Println("parent", lastRm.parent, "root", root)
			lastRm.parent = append(lastRm.parent, root)

		}

		countofparents := len(lastRm.parent)
		lastRm.parent[countofparents-1].children = append(lastRm.parent[countofparents-1].children, lastRm)
		// fmt.Println("endrm", *lastRm)

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
		//  fmt.Println("firstroom", *firstRm)
	} else {
		fmt.Println("___________________________________________________")
		fmt.Printf("constructing other room %s...\n", rmToAddName)
		test = nil
		test = append(test, root)

		if dup(checkdup) != rmToAddName {
			roomToAdd = &room{
				parent:   test,
				name:     rmToAddName,
				occupied: false,
			}
		} else {
			roomToAdd = &room{
				parent:   append(roomToAdd.parent, root),
				name:     rmToAddName,
				occupied: false,
			}
		}

		countofparents := len(roomToAdd.parent)
		roomToAdd.parent[countofparents-1].children = append(roomToAdd.parent[countofparents-1].children, roomToAdd)
		fmt.Println("parent.children", *roomToAdd.parent[countofparents-1].children[0])
		findChildren(roomToAdd, rmToAddName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
		// fmt.Println("roomtoadd", *roomToAdd)
	}
	Farm = append(Farm, *roomToAdd)
	// fmt.Println("firstrm", firstRm)
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
func dup(s []string) string {
	duplicate := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		if duplicate[s[i]] == true {
			return s[i]
		} else {
			duplicate[s[i]] = true
		}
	}
	return ""
}
