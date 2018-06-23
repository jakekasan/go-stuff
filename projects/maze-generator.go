package main

import (
	"fmt"
	"math/rand"
)

type Cell struct {
	x       int
	y       int
	visited bool
	walls   []bool
}

type Pos struct {
	x int
	y int
}

const SIZE = 100 // the number of columns/rows in the maze. The actual size will be SIZE*SIZE

func getOptions(path *[]Pos, maze *[SIZE][SIZE]Cell) []Pos {
	// look at options
	options := make([]Pos, 0)
	if (*path)[len((*path))-1].x > 0 {
		if (*maze)[(*path)[len((*path))-1].y][(*path)[len((*path))-1].x-1].visited == false {
			options = append(options, Pos{x: (*path)[len((*path))-1].x - 1, y: (*path)[len((*path))-1].y})
		}
	}
	if (*path)[len((*path))-1].x < (len((*maze)[0]) - 1) {
		if (*maze)[(*path)[len((*path))-1].y][(*path)[len((*path))-1].x+1].visited == false {
			options = append(options, Pos{x: (*path)[len((*path))-1].x + 1, y: (*path)[len((*path))-1].y})
		}
	}
	if (*path)[len((*path))-1].y > 0 {
		if (*maze)[(*path)[len((*path))-1].y-1][(*path)[len((*path))-1].x].visited == false {
			options = append(options, Pos{x: (*path)[len((*path))-1].x, y: (*path)[len((*path))-1].y - 1})
		}
	}
	if (*path)[len((*path))-1].y < (len((*maze)) - 1) {
		fmt.Println((*path)[len((*path))-1].x)
		if (*maze)[(*path)[len((*path))-1].y+1][(*path)[len((*path))-1].x].visited == false {
			options = append(options, Pos{x: (*path)[len((*path))-1].x, y: (*path)[len((*path))-1].y + 1})
		}
	}
	return options
}

func moveTo(choice Pos, maze *[SIZE][SIZE]Cell, path *[]Pos) {
	*path = append(*path, choice)
	(*maze)[(*path)[len((*path))-1].y][(*path)[len((*path))-1].x].visited = true
}

func main() {
	fmt.Println("Hello world!")

	var maze [SIZE][SIZE]Cell

	for i, _ := range maze {
		for j, _ := range maze[i] {
			maze[i][j] = Cell{x: j, y: i, visited: false, walls: []bool{true, true, true, true}}
		}
	}

	// set starting position: 0,0

	path := make([]Pos, 0)

	start := Pos{x: 0, y: 0}

	path = append(path, start)

	for len(path) > 0 {
		// get options
		//fmt.Println("X:", path[len(path)-1].x, "\tY:", path[len(path)-1].y)
		options := getOptions(&path, &maze)
		if len(options) < 1 {
			path = path[:(len(path) - 1)]
		} else {
			choice := options[(rand.Int() % len(options))]
			moveTo(choice, &maze, &path)
		}

	}
}
