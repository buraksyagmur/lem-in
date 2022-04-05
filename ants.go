package main

import "fmt"

type ant struct {
	id        int
	curRoom   *room
	pathOfAnt []*room
}

var (
	ants                []ant
	roomspassed         int
	Combinations        [][]*room
	allways             [][]*room
	way                 []*room
	way2                []*room
	cleanway            [][]*room
	childrenOfFirstRoom int
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

func FindingPath(currentRoom *room) [][]*room {
	if len(currentRoom.children) != 0 {
		for i := 0; i < len(currentRoom.children); i++ {
			fmt.Println("number of i", i, "room", *currentRoom, "child", *currentRoom.children[i])
			way2 = append(way2, currentRoom.children[i])
			FindingPath(currentRoom.children[i])
		}
	} else {

		allways = append(allways, way2)
		way2 = nil
	}

	return allways
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

	for i := 0; i < len(*paths); i++ {
		if (*paths)[i] == nil {
			*paths = append((*paths)[:i], (*paths)[i+1:]...)
		}
	}
}

func SortPaths(ways [][]*room) [][]*room {
	for i := 0; i < len(ways)-1; i++ {
		if len(ways[i]) > len(ways[i+1]) {
			ways[i], ways[i+1] = ways[i+1], ways[i]
		}
	}

	for k := 0; k < len(ways)-1; k++ {
		if len(ways[len(ways)-1]) < len(ways[k]) {
			ways[len(ways)-1], ways[k] = ways[k], ways[len(ways)-1]
		}
	}
	return ways
}

// func FindBestCombinations (ways [][]*room , countOfAnts int) [][]*room{

// return ways
// }

func ClearPath(ways [][]*room) [][]*room {
	childrenOfFirstRoom = len(firstRm.children)
	var countofloop int
	for i := countofloop + 1; i < len(ways); i++ {
		for k := 0; k < len(ways[countofloop])-1; k++ {
			for t := 0; t < len(ways[i])-1; t++ {
				if ways[countofloop][k].name == ways[i][t].name {
				}
			}
		}
	}
	return ways
}
