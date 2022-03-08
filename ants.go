package main

import "fmt"

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
