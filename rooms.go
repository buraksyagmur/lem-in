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
	Farm2    []*room
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
		if startRmName == destConnRmNames[c] {
			beginConnRmNames[c], destConnRmNames[c] = destConnRmNames[c], beginConnRmNames[c]
		}
		if rmToAddName != startRmName {
			for m := 0; m < len(firstRm.children); m++ {
				if firstRm.children[m].name == destConnRmNames[c] && startRmName != beginConnRmNames[c] {
					beginConnRmNames[c], destConnRmNames[c] = destConnRmNames[c], beginConnRmNames[c]
				}
			}
		}
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
			// fmt.Println("infos", roomToAdd, destRmName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
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
		if len(dup(checkdup)) != 0 {
			for i := 0; i < len(dup(checkdup)); i++ {
				if (dup(checkdup)[i]) == rmToAddName {
					lastRm.parent = append(lastRm.parent, root)
				} else {
					lastRm = &room{
						parent:   test,
						name:     rmToAddName,
						children: nil,
						occupied: false,
					}
				}
			}
		} else {
			lastRm = &room{
				parent:   test,
				name:     rmToAddName,
				children: nil,
				occupied: false,
			}
			// fmt.Println("parent", lastRm.parent, "root", root)
		}
		countofparents := len(lastRm.parent)
		lastRm.parent[countofparents-1].children = append(lastRm.parent[countofparents-1].children, lastRm)
		// fmt.Println("lastone", *lastRm.parent[countofparents-1], "count", countofparents)
		// fmt.Println("endrm", *lastRm)

		return lastRm
	} else if rmToAddName == startRmName { // start Room special case
		fmt.Println("___________________________________________________")
		fmt.Printf("constructing start room %s...\n", rmToAddName)
		firstRm = &room{
			parent:   nil,
			name:     rmToAddName,
			occupied: true,
		}
		findChildren(firstRm, rmToAddName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
		// firstRm = roomToAdd
		//  fmt.Println("firstroom", *firstRm)
		Farm = append(Farm, *firstRm)
	} else {
		fmt.Println("___________________________________________________")
		fmt.Printf("constructing other room %s...\n", rmToAddName)
		test = nil
		test = append(test, root)
		if len(dup(checkdup)) != 0 {
			for i := 0; i < len(dup(checkdup)); i++ {
				if (dup(checkdup)[i]) == rmToAddName {
					// fmt.Println("roomtoaddparent", roomToAdd, "Farm", Farm)
					for t := 0; t < len(Farm); t++ {
						if Farm[t].name == rmToAddName {
							roomToAdd = &Farm[t]
							roomToAdd.parent = append(roomToAdd.parent, root)
							// fmt.Println("roomtoaddparent", roomToAdd)
						}
					}
				} else {
					roomToAdd = &room{
						parent:   test,
						name:     rmToAddName,
						occupied: false,
					}
				}
			}
		} else {
			roomToAdd = &room{
				parent:   test,
				name:     rmToAddName,
				occupied: false,
			}
		}

		countofparents := len(roomToAdd.parent)
		for r := 0; r < len(roomToAdd.parent[countofparents-1].children); r++ {
			if roomToAdd.parent[countofparents-1].children[r].name == roomToAdd.name {
				roomToAdd.parent[countofparents-1].children[r] = roomToAdd
			} else {
				roomToAdd.parent[countofparents-1].children = append(roomToAdd.parent[countofparents-1].children, roomToAdd)
			}
		}

		// fmt.Println("parent.children", *roomToAdd.parent[countofparents-1].children[0])
		// fmt.Println(roomToAdd, rmToAddName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
		findChildren(roomToAdd, rmToAddName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
		// fmt.Println("roomtoadd", *roomToAdd)
		// for p := 0; p < len(Farm); p++ {
		// 	if Farm[p].name == roomToAdd.name {

		// 		fmt.Println("old", Farm)
		// 		RemoveElement(Farm, p)
		// 		fmt.Println("new", Farm)

		// }
		Farm = CheckFarmDup(Farm, roomToAdd.name)
		fmt.Println("roomtoaddname", roomToAdd.name)
		Farm = append(Farm, *roomToAdd)
	}

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
func dup(s []string) []string {
	var result []string
	duplicate := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		if duplicate[s[i]] == true {
			result = append(result, s[i])
		} else {
			duplicate[s[i]] = true
		}
	}
	return result
}

func RemoveElement(s []room, i int) []room {
	return append(s[:i], s[i+1:]...)
}

func CheckFarmDup(s []room, t string) []room {
	for k := 0; k < len(s); k++ {
		if s[k].name == t {
			// fmt.Println("oldname", s[k].name, "roomtoadd", t)
			s = RemoveElement(s, k)
			// fmt.Println("newname", s[k].name, "roomtoadd", t)
			CheckFarmDup(s, t)
		}
	}

	return s
}
