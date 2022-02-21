package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	antnumber  string
	startline  int
	endline    int
	endroom    string
	slccontent []string
)

func readnote(textfile string) {
	// f, err := os.Open("examples/" + textfile)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	// scanner := bufio.NewScanner(f)
	// var line int
	// var line2 int
	// var line3 int
	// var line4 int
	// for scanner.Scan() {
	// 	if line == 0 {
	// 		antnumber = scanner.Text()
	// 	}
	// 	line++
	// 	if strings.Contains(scanner.Text(), "##start") {
	// 		startline = line2+1
	// 	}
	// 	line2++

	// 	if strings.Contains(scanner.Text(), "##end") {
	// 		endline = line3+1
	// 	}
	// 	line3++

	// 	if line == endline+1 {
	// 		endroom = scanner.Text()
	// 	}
	// 	line4++
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatalln(err)
	// }

	content, err := ioutil.ReadFile("examples/" + textfile)
	if err != nil {
		log.Fatal(err)
	}

	slccontent = strings.Split(string(content), "\n")
	for i := 0; i < len(slccontent); i++ {
		if slccontent[i] == "##start" {
			startline = i
		}
		if slccontent[i] == "##end" {
			endline = i
		}
	}
}

func main() {
	readnote(os.Args[1])
	fmt.Println("ant number:", slccontent[0])

	fmt.Println("startline:", startline)
	fmt.Println("startroom:", slccontent[startline+1])
	for i := 3; i < endline; i++ {
		fmt.Println("rooms:", slccontent[i])
	}
	fmt.Println("endline:", endline)
	fmt.Println("endroom:", slccontent[endline+1])
	for k := endline + 2; k < len(slccontent); k++ {
		fmt.Println("connections:", slccontent[k])
	}
}
