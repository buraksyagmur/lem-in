package main

import "fmt"

type ant struct {
	id      int
	curRoom *room
}

var (
	ants         []ant
	roomspassed  int
	Combinations [][]*room
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

// // have to consider the first ant in to the end room
// func (a *ant) move(out, in *room) {
// 	if !in.occupied {
// 		out.occupied = false
// 		in.occupied = true
// 		a.curRoom = in
// 	}
// }

// // have to consider the last ant out
// func (a *ant) startMove(in *room) {
// 	if !in.occupied {
// 		in.occupied = true
// 		a.curRoom = in
// 	}
// }

// func (a *ant) checkState(rm *room) {
// 	if rm != lastRm {
// 	}
// }
func SwapFarm(Farm []room) {
	for i := 0; i < len(Farm)/2; i++ {
		Farm[i], Farm[(len(Farm)-1)-i] = Farm[(len(Farm)-1)-i], Farm[i]
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

func AllPaths(Farm []*room, currentRoom *room, Sroom *room, Eroom *room) [][]*room {
	var Paths [][]*room
	if currentRoom.name == lastRm.name {
		for i := 0; i < (len(Paths)); i++ {
			for k := 0; k < len(Paths[i]); k++ {
				for t := 0; t < len(Farm); t++ {
					if Paths[i][k] == Farm[t] {
						continue
					}
				}
			}
		}
	}
	for p := 0; p < (len(Paths)); p++ {
		for r := 0; r < len(Sroom.children); r++ {
		}
	}

	return Paths
}

func FindAllPossiblePaths(path []*room, currentRoom room, paths *[][]*room, previousRoom *room) {
	if currentRoom.name == lastRm.name {
		var skipPath bool
		for i := 0; i < len(path); i++ {
			if path[i].name == firstRm.name {
				skipPath = true
				break
			}
		}

		if len(*paths) == 0 {
			*paths = append((*paths), nil)
		} else if (*paths)[len(*paths)-1] != nil {
			*paths = append((*paths), nil)
		}

		for i := 0; i < len(path); i++ {
			if !skipPath {
				(*paths)[len(*paths)-1] = append((*paths)[len(*paths)-1], path[i])
			} else {
				break
			}
		}
	}

	for i := 0; i < len(currentRoom.children); i++ {
		var toContinue bool

		for k := 0; k < len(path); k++ {
			if path[k].name == currentRoom.children[i].name {
				toContinue = true
				break
			}
		}

		if !toContinue {
			pathToPass := path
			pathToPass = append(pathToPass, currentRoom.children[i])
			FindAllPossiblePaths(pathToPass, *currentRoom.children[i], paths, &currentRoom)
			pathToPass = path
		}
	}

	if paths != nil {
		for i := 0; i < len(*paths); i++ {
			if (*paths)[i] == nil {
				*paths = append((*paths)[:i], (*paths)[i+1:]...)
			}
		}
	}
}
