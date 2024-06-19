package main

import (
	"fmt"
	"math/rand"
	"time"
)

const ROWS = 100
const COLS = 100

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func cursorTopLeft() {
    fmt.Print("\033[H")
}

func countNeighbour(world *[100][100]string, i int, j int) (int, int) {
	liveNeighbour := 0
	deadNeighbour := 0

	if world[i][j] == "#" {
		liveNeighbour = 1
	} else {
		deadNeighbour = 1
	}

	return liveNeighbour, deadNeighbour
}

/*
- - -
- 0 -
- - -
check all neighbours
*/
func countNeighbours(world *[100][100]string, i int, j int) (int, int) {
	liveNeighbours := 0
	deadNeighbours := 0

	topLeft_live, topLeft_dead := 0, 0
	top_live, top_dead := 0, 0
	topRight_live, topRight_dead := 0, 0
	left_live, left_dead := 0, 0
	right_live, right_dead := 0, 0
	bottomLeft_live, bottomLeft_dead := 0, 0
	bottom_live, bottom_dead := 0, 0
	bottomRight_live, bottomRight_dead := 0, 0

	i_minus_1, i_plus_1, j_minus_1, j_plus_1 := 0, 0, 0, 0

	if i == 0 {
		i_minus_1 = 99
	} else {
		i_minus_1 = i - 1
	}

	if i == 99 {
		i_plus_1 = 0
	} else {
		i_plus_1 = i + 1
	}

	if j == 0 {
		j_minus_1 = 99
	} else {
		j_minus_1 = j - 1
	}

	if j == 99 {
		j_plus_1 = 0
	} else {
		j_plus_1 = j + 1
	}

	topLeft_live, topLeft_dead = countNeighbour(world, i_minus_1, j_minus_1)
	left_live, left_dead = countNeighbour(world, i, j_minus_1)
	bottomLeft_live, bottomLeft_dead = countNeighbour(world, i_plus_1, j_minus_1)
	top_live, top_dead = countNeighbour(world, i_minus_1, j)
	bottom_live, bottom_dead = countNeighbour(world, i_plus_1, j)
	topRight_live, topRight_dead = countNeighbour(world, i_minus_1, j_plus_1)
	right_live, right_dead = countNeighbour(world, i, j_plus_1)
	bottomRight_live, bottomLeft_dead = countNeighbour(world, i_plus_1, j_plus_1)

	liveNeighbours = topLeft_live + left_live + bottomLeft_live + topRight_live + right_live + bottomRight_live + top_live + bottom_live
	deadNeighbours = topLeft_dead + left_dead + bottomLeft_dead + topRight_dead + right_dead + bottomRight_dead + top_dead + bottom_dead

	return liveNeighbours, deadNeighbours
}
func generateWorld(world *[100][100]string) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if rand.Intn(5) == 4 {
				world[i][j] = "#"
            }
		}
	}
}

func updateWorldState(world *[100][100]string) {
	// * -> alive, "" -> unpopulated, # -> dead
	// Question -> Should the states change at the same time? Yes, probably.
	// Doubt -> Do I need to initialise next with initial values of world every time?
	next := [100][100]string{}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			liveNeighboursVal, _ := countNeighbours(world, i, j)
			if world[i][j] == "#" {
				if liveNeighboursVal < 2 {
					next[i][j] = ""
				} else if liveNeighboursVal >= 2 && liveNeighboursVal <= 3 {
					next[i][j] = "#"
				} else {
					next[i][j] = ""
				}
			} else {
				if liveNeighboursVal == 3 {
					next[i][j] = "#"
				}
			}
		}
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			world[i][j] = next[i][j]
		}
	}
}

func printWorld(world *[100][100]string) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			//fmt.Printf("Hello \n");
            if world[i][j] == "" {
                fmt.Printf(" ")
            } else {
                fmt.Printf("%s", world[i][j])
            }
		}
		fmt.Printf("\n")
	}
}

//func main() {
//	world := [100][100]string{}
//    generateWorld(&world)
//        printWorld(&world)
//}

func main() {
    cursorTopLeft()
    clearScreen()
	world := [100][100]string{}
	generateWorld(&world)
	for i := 0; i < 5000; i++ {
        //cursorTopLeft()
        clearScreen()
		printWorld(&world)
		//clrscr()
        time.Sleep(24 * time.Millisecond)
        updateWorldState(&world)
	}
}
