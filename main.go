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
	RoomsandConnections []string
)

type room struct {
	name     string
	parent   *room
	children []room
	occupied bool
}

// split ant to another file later
type ant struct {
	id      int
	curRoom *room
}

func readnote(textfile string) {
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
	RoomsandConnections = strings.Split(string(content), "\n")
	RoomsandConnections = append(RoomsandConnections[:endline], RoomsandConnections[endline+2:]...)

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

func main() {
	readnote(os.Args[1])
	fmt.Println("ant number:", slccontent[0])
	fmt.Println("startroom:", slccontent[startline+1])
	for i := 3; i < connectionStartLine; i++ {
		fmt.Println("other rooms:", RoomsandConnections[i])
	}
	fmt.Println("endline:", endline)
	fmt.Println("endroom:", slccontent[endline+1])
	for m := connectionStartLine; m < len(RoomsandConnections); m++ {
		fmt.Println("connections:", RoomsandConnections[m])
	}

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

	numOfAnt, _ := strconv.Atoi(slccontent[0])
	ants := make([]ant, numOfAnt)
	fmt.Print(len(ants))
}
