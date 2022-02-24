package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	antnumber           string
	startline           int
	endline             int
	endroom             string
	slccontent          []string
	connectionStartLine int
	RoomsandConnections []string
)

type room struct {
	name      string
	parent    *room
	leftChild *room
	rightSib  *room
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

func createRoomNode(name string, p, l, r room) {
	var newRoom room
	newRoom.parent = p
	newRoom.leftChild = l
	newRoom.rightSib = r
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
}
