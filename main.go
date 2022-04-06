package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	startline           int
	endline             int
	slccontent          []string
	connectionStartLine int
	// RoomsandConnections []string
	numOfAnt      int
	nameofStart   string
	nameofEnd     string
	numberofRooms int
	antFarmRooms  *room
)

func readnote(textfile string) []string {
	content, err := ioutil.ReadFile("examples/" + textfile)
	if err != nil {
		log.Fatal(err)
	}

	slccontent = strings.Split(string(content), "\n")
	for l := 0; l < len(slccontent); l++ {
		for t := 0; t < len(slccontent[l]); t++ {
			if string(slccontent[l][0]) == "#" {
				if string(slccontent[l][1]) != "#" {
					slccontent = append(slccontent[:l], slccontent[l+1:]...)
				}
				for i := 0; i < len(slccontent); i++ {
					if slccontent[i] == "##start" {
						startline = i
					}
					if slccontent[i] == "##end" {
						endline = i
					}
				}
			}
		}
	}
	RoomsandConnections := strings.Split(string(content), "\n")
	// RoomsandConnections = append(RoomsandConnections[3:endline], RoomsandConnections[endline+2:]...)
	// including start and end room
	RoomsandConnections = append(RoomsandConnections[2:endline], RoomsandConnections[endline+1:]...)
	// fmt.Println(RoomsandConnections)
	for l := 0; l < len(RoomsandConnections); l++ {
		for t := 0; t < len(RoomsandConnections[l]); t++ {
			if string(RoomsandConnections[l][0]) == "#" {
				if string(RoomsandConnections[l][1]) != "#" {
					RoomsandConnections = append(RoomsandConnections[:l], RoomsandConnections[l+1:]...)
				}
			}
		}
	}
	for m := len(RoomsandConnections) - 1; m >= 0; m-- {
		for k := 0; k < len(RoomsandConnections[m]); k++ {
			if string(RoomsandConnections[m][k]) == "-" {
				connectionStartLine = m
				break
			}
		}
	}
	stringfirstname := slccontent[startline+1]
	slicefirstname := strings.Split(stringfirstname, " ")
	nameofStart = slicefirstname[0]
	stringlastname := slccontent[endline+1]
	slicelastname := strings.Split(stringlastname, " ")
	nameofEnd = slicelastname[0]
	return RoomsandConnections
}

func main() {
	roomsandConnections := readnote(os.Args[1])
	numOfAnt, _ = strconv.Atoi(slccontent[0])
	// fmt.Println("ant number:", slccontent[0])
	// fmt.Println("startroom:", slccontent[startline+1])
	// for i := 1; i < connectionStartLine; i++ { // exclude the start room
	// 	fmt.Println("other rooms:", RoomsandConnections[i])
	// }
	// fmt.Println("endline:", endline)
	// fmt.Println("endroom:", slccontent[endline+1])
	// for m := connectionStartLine; m < len(RoomsandConnections); m++ {
	// 	fmt.Println("connections:", RoomsandConnections[m])
	// }

	// constructing rooms
	antFarmRooms = Rooms(roomsandConnections)
	// fmt.Println("antFarmRooms", antFarmRooms)
	// fmt.Println(lastRm, "lastroom")
	// fmt.Println("Farm1", Farm[0])
	// fmt.Println("Farm2", Farm[1])
	// fmt.Println("Farm", Farm)
	// fmt.Println("Farm1", Farm)

	// ShortestPath(Farm[numOfAnt])
	SwapFarm(Farm)
	Farm = append(Farm, *lastRm)
	numberofRooms = len(Farm)
	for i := 0; i < len(Farm); i++ {
		fmt.Println(i, Farm[i])
	}

	// fmt.Println("allpath", FindingPath(firstRm))
	fmt.Println("numberofways", len(allways))
	var firstchildrennames []string
	for l := 0; l < len(firstRm.children); l++ {
		firstchildrennames = append(firstchildrennames, firstRm.children[l].name)
	}
	// var allways2 [][]*room
	// for y := 0; y < len(allways); y++ {
	// 	if allways[y][len(allways[y])-1].name == lastRm.name {
	// 		for e := 0; e < len(firstchildrennames); e++ {
	// 			if allways[y][0].name == firstchildrennames[e] {
	// 				allways2 = append(allways2, allways[y])
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println("correctwaysnumber", len(allways2))
	// for o := 0; o < len(allways2[0]); o++ {
	// 	fmt.Println("1allcorrectways1", allways2[0][o])
	// }
	// for o := 0; o < len(allways2[1]); o++ {
	// 	fmt.Println("2allcorrectways2", allways2[1][o])
	// }
	// for o := 0; o < len(allways2[2]); o++ {
	// 	fmt.Println("3allcorrectways3", allways2[2][o])
	// }
	// for o := 0; o < len(allways2[3]); o++ {
	// 	fmt.Println("4allcorrectways4", allways2[3][o])
	// }
	// antfarm := CreatingAnts()
	// walk(antfarm)
	// FindAllPossiblePaths(make([]*room, 0), Farm[0], &Combinations, &Farm[0])
	// for k := 0; k < len(Combinations); k++ {
	// 	fmt.Println("id ", k, "combinations", Combinations[k])
	// }
	// ants
	// ants := make([]ant, numOfAnt)
	// fmt.Println("no. of ants: ", len(ants))

	var allPaths [][]*room
	FindAllPossiblePaths(make([]*room, 0), Farm[0], &allPaths, &Farm[0])
	for i := 0; i < len(allPaths); i++ {
		allPaths = SortPaths(allPaths)
	}
	for i := 0; i < len(allPaths); i++ {
		for k := 0; k < len(allPaths[i]); k++ {
			fmt.Println(i, "thats the all sorted paths ", allPaths[i][k])
		}
	}

	cleanway = ClearPath(allPaths)

	// for i := 0; i < len(cleanway); i++ {
	// 	for k := 0; k < len(cleanway[i]); k++ {
	// 		fmt.Println(i, "thats the all best paths ", cleanway[i][k])
	// 	}
	// }

	var anotherbool2 bool = false
	for i := 0; i < len(appendWays); i++ {
		for k := 0; k < len(allPaths[len(allPaths)-1]); k++ {
			if allPaths[len(allPaths)-1][k] == appendWays[i] {
				anotherbool2 = true
			}
		}
	}
	if !anotherbool2 {
		cleanway = append(CombinatedRooms, allPaths[len(allPaths)-1])
	}
	for i := 0; i < len(cleanway); i++ {
		for p := 0; p < len(cleanway[i]); p++ {
			fmt.Println(i, "cleanway", cleanway[i][p])
		}
	}
}
