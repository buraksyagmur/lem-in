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
	numOfAnt int
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
	antFarmRooms := Rooms(roomsandConnections)
	fmt.Println("antFarmRooms", antFarmRooms)
	CreatingAnts()
	// ants
	// ants := make([]ant, numOfAnt)
	// fmt.Println("no. of ants: ", len(ants))
}
