package main

import "fmt"

type ant struct {
	id      int
	curRoom *room
}

var (
	ants        []ant
	roomspassed int
)

func CreatingAnts() []ant {
	if firstRm != nil {
		for i := 1; i < numOfAnt+1; i++ {
			var allants ant

			allants.id = i
			allants.curRoom = firstRm

			ants = append(ants, allants)
		}
	}
	// fmt.Println("ants", ants)
	return ants
}

// have to consider the first ant in to the end room
func (a *ant) move(out, in *room) {
	if !in.occupied {
		out.occupied = false
		in.occupied = true
		a.curRoom = in
	}
}

// have to consider the last ant out
func (a *ant) startMove(in *room) {
	if !in.occupied {
		in.occupied = true
		a.curRoom = in
	}
}

func (a *ant) checkState(rm *room) {
	if rm != lastRm {
	}
}

func walk(antfarm []ant) {
	var allpassed bool = true
	for i := 0; i < len(antfarm); i++ {

		if antfarm[i].curRoom.name == lastRm.name {
			continue
		}
		if antfarm[i].curRoom.children[0].occupied {
			continue
		}

		antfarm[i].curRoom.occupied = false
		antfarm[i].curRoom = antfarm[i].curRoom.children[0]
		if antfarm[i].curRoom.name != lastRm.name {
			antfarm[i].curRoom.occupied = true
		}
		allpassed = false

		fmt.Print("L", antfarm[i].id, "-", antfarm[i].curRoom.name, " ")

	}
	if allpassed {
		return
	} else {
		fmt.Println(" ")
		walk(antfarm)
	}
}

func ShortestPath(rm *room) {
	// fmt.Println("roomname", rm)
	// fmt.Println("childname", rm.children[0].name)
	// fmt.Println("endroomname", lastRm.name)

	if rm.name != lastRm.name {
		roomspassed++
		for i := 0; i < len(rm.children); i++ {
			fmt.Println(i)
			ShortestPath(rm.children[i])
		}
	} else {
		fmt.Println(roomspassed)
		fmt.Println(rm.name)
	}
}
