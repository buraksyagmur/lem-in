package main

import "fmt"

type ant struct {
	id      int
	curRoom *room
}

var ants []ant

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
	for i := 0; i < len(antfarm)-1; i++ {

		// if !antfarm[i].curRoom.children[0].occupied {
		// 	continue
		// }

		// antfarm[i].curRoom.occupied = false
		antfarm[i].curRoom = antfarm[i].curRoom.children[0]
		antfarm[i].curRoom.occupied = true
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
