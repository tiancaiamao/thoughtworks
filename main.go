package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
	g, err := MakeGraph(inputGraph)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	lines := bytes.Split(inputQuestion, []byte{'\n'})
	for _, line := range lines {
		Solve(line, g)
	}
}

func MakeGraph(input []byte) (*Graph, error) {
	err := errors.New("I have no idea what you are talking about")
	if !bytes.HasPrefix(input, []byte("Graph:")) {
		return nil, err
	}

	g := NewGraph()
	input = input[6:]
	for len(input) >= 3 {
		// trim space
		for input[0] == ' ' {
			input = input[1:]
		}

		a := input[0]
		b := input[1]
		if a < 'A' || a > 'Z' || b < 'A' || b > 'Z' {
			return nil, err
		}

		input = input[2:]
		idx := bytes.IndexFunc(input, func(r rune) bool {
			if r < '0' || r > '9' {
				return true
			}
			return false
		})

		var c int
		var err error
		if idx == -1 {
			c, err = strconv.Atoi(string(input))
		} else {
			c, err = strconv.Atoi(string(input[:idx]))
		}

		if err != nil {
			return nil, err
		}
		g.AddEdge(a, b, uint(c))

		idx = bytes.IndexByte(input, ',')
		if idx > 0 {
			input = input[idx+1:]
		} else {
			break
		}
	}
	return g, nil
}

func Solve(line []byte, g *Graph) {
	quest1 := []byte("The distance of the route ")
	quest2 := []byte("The number of trips starting at ")
	quest3 := []byte("The length of the shortest route (in terms of distance to travel) from ")
	quest4 := []byte("The number of different routes from ")
	idx := bytes.Index(line, []byte{'.'})
	if idx == -1 {
		fmt.Fprintf(os.Stderr, "I have no idea what you are talking about")
		return
	}

	num, err := strconv.Atoi(string(line[:idx]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "I have no idea what you are talking about")
		return
	}

	line = bytes.TrimSpace(line[idx+1:])

	if bytes.HasPrefix(line, quest1) {
		line = line[len(quest1):]
		input, err := ParseQuest1(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return
		}

		SolveQuest1(input, num, g)
	} else if bytes.HasPrefix(line, quest2) {
		line = line[len(quest2):]
		input, err := ParseQuest2(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return
		}
		SolveQuest2(input, num, g)
	} else if bytes.HasPrefix(line, quest3) {
		line = line[len(quest3):]
		from, to, err := ParseQuest3(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return
		}
		SolveQuest3(from, to, num, g)
	} else if bytes.HasPrefix(line, quest4) {
		line = line[len(quest4):]
		input, err := ParseQuest4(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return
		}
		SolveQuest4(input, num, g)
	} else {
		fmt.Fprintf(os.Stderr, "I have no idea what you are talking about")
	}
}

type Quest4 struct {
	from byte
	to   byte
	dist int
}

// example input: "C to C with a distance of less than 30."
func ParseQuest4(line []byte) (ret *Quest4, err error) {
	err = errors.New("I have no idea what you are talking about")
	ret = new(Quest4)
	ret.from = line[0]
	line = line[1:]

	str := " to "
	if !bytes.HasPrefix(line, []byte(str)) {
		return
	}

	line = line[len(str):]
	ret.to = line[0]
	line = line[1:]

	str = " with a distance of less than "
	if !bytes.HasPrefix(line, []byte(str)) {
		return
	}
	line = line[len(str):]

	idx := bytes.IndexByte(line, '.')
	if idx != len(line)-1 {
		return
	}

	ret.dist, err = strconv.Atoi(string(line[:idx]))
	return
}

func SolveQuest4(input *Quest4, num int, g *Graph) {
	n := BruteForceDistance(g, input.dist, input.from, input.to)
	fmt.Printf("Output #%d: %d\n", num, n)
}

// example input: "A to B."
func ParseQuest3(line []byte) (from byte, to byte, err error) {
	err = errors.New("I have no idea what you are talking about")
	str := "* to *."
	if len(line) != len(str) {
		return
	}

	for i := 0; i < len(str); i++ {
		if str[i] != '*' && str[i] != line[i] {
			return
		}
	}

	from = line[0]
	to = line[5]
	err = nil
	return
}

func SolveQuest3(from byte, to byte, num int, g *Graph) {
	dist := g.Dijkstra(from, to)
	if dist == UINT_MAX {
		fmt.Printf("Output #%d: NO SUCH ROUTE\n", num)
	} else {
		fmt.Printf("Output #%d: %d\n", num, dist)
	}
}

type Quest2 struct {
	start   byte
	end     byte
	exactly bool // exactly or maximum
	stops   int
}

// example input: "C and ending at C with a maximum of 3 stops."
func ParseQuest2(line []byte) (*Quest2, error) {
	ret := new(Quest2)
	err := errors.New("I have no idea what you are talking about")

	// parse start
	if line[0] < 'A' || line[0] > 'Z' {
		return nil, err
	}
	ret.start = line[0]
	line = line[1:]
	str := " and ending at "
	if !bytes.HasPrefix(line, []byte(str)) {
		return nil, err
	}
	line = line[len(str):]

	// parse end
	if line[0] < 'A' || line[0] > 'Z' {
		return nil, err
	}
	ret.end = line[0]
	line = line[1:]
	str = " with "

	if !bytes.HasPrefix(line, []byte(str)) {
		return nil, err
	}
	line = line[len(str):]

	// parse type
	maximum := "a maximum of "
	exactly := "exactly "
	if bytes.HasPrefix(line, []byte(maximum)) {
		ret.exactly = false
		line = line[len(maximum):]
	} else if bytes.HasPrefix(line, []byte(exactly)) {
		ret.exactly = true
		line = line[len(exactly):]
	} else {
		return nil, err
	}

	// parse stops
	idx := bytes.Index(line, []byte(" stops."))
	if idx == -1 {
		return nil, err
	}
	ret.stops, err = strconv.Atoi(string(line[:idx]))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func SolveQuest2(input *Quest2, num int, g *Graph) {
	within, exactly := BruteForceNStops(g, input.stops, input.start, input.end)
	if input.exactly {
		fmt.Printf("Output #%d: %d\n", num, exactly)
	} else {
		fmt.Printf("Output #%d: %d\n", num, within)
	}
}

func SolveQuest1(input []byte, num int, g *Graph) {
	start := input[0]
	current, ok := g.Name2Vertex[start]
	if !ok {
		fmt.Printf("I have no idea what you are talking about")
		return
	}

	result := uint(0)
	for i := 1; i < len(input); i++ {
		name := input[i]
		next, ok := g.Name2Vertex[name]
		if !ok {
			fmt.Printf("I have no idea what you are talking about")
			return
		}

		find := false
		for e := g.Edges[current]; e != nil; e = e.Next {
			if e.ToVertex == next {
				find = true
				result += e.Weight
				current = next
				break
			}
		}

		if !find {
			fmt.Printf("Output #%d: NO SUCH ROUTE\n", num)
			return
		}
	}
	fmt.Printf("Output #%d: %d\n", num, result)
	return
}

// line format similar to A-B-C.
func ParseQuest1(line []byte) ([]byte, error) {
	err := errors.New("I have no idea what you are talking about")
	input := make([]byte, 0, 5)
	for i := 0; i < len(line); {
		if line[i] < 'A' || line[i] > 'Z' {
			return input, err
		}

		input = append(input, line[i])
		i++
		if i >= len(line) {
			return input, err
		}

		if line[i] == '-' {
			i++
		} else if line[i] == '.' {
			break
		} else {
			return input, err
		}
	}
	return input, nil
}
