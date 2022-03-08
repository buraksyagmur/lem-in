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
	fmt.Println("ants", ants)
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
	if rm != endRoom {

	}
}
