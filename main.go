package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const inputGraph string = `Graph: AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7`

const inputQuestion string = `1. The distance of the route A-B-C.
2. The distance of the route A-D.
3. The distance of the route A-D-C.
4. The distance of the route A-E-B-C-D.
5. The distance of the route A-E-D.
6. The number of trips starting at C and ending at C with a maximum of 3 stops.
7. The number of trips starting at A and ending at C with exactly 4 stops.
8. The length of the shortest route (in terms of distance to travel) from A to C.
9. The length of the shortest route (in terms of distance to travel) from B to B.
10.The number of different routes from C to C with a distance of less than 30.`

func usage() {
	fmt.Fprintf(os.Stderr, "usage: trains graphFile questionFile")
	os.Exit(-1)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	inputGraph, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		usage()
	}
	inputQuestion, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		usage()
	}

	Process(inputGraph, inputQuestion)
	return
}

func Process(inputGraph, inputQuestion []byte) {
	g, err := ParseGraph(inputGraph)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	lines := bytes.Split(inputQuestion, []byte{'\n'})
	for _, line := range lines {
		prob, err := ParseQuestion(line)
		if err != nil {
			fmt.Println("I have no idea what you are talking about")
		} else {
			prob.Solve(g)
		}
	}
}
