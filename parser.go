package main

import (
	"bytes"
	"errors"
	"strconv"
)

type Dispatch struct {
	prefix []byte
	parser func(line []byte, id int) (Problem, error)
}

var dispatchs []Dispatch
var errFormat error

func init() {
	dispatchs = make([]Dispatch, 4)
	dispatchs[0] = Dispatch{
		prefix: []byte("The distance of the route "),
		parser: parseQuest1,
	}
	dispatchs[1] = Dispatch{
		prefix: []byte("The number of trips starting at "),
		parser: parseQuest2,
	}
	dispatchs[2] = Dispatch{
		prefix: []byte("The length of the shortest route (in terms of distance to travel) from "),
		parser: parseQuest3,
	}
	dispatchs[3] = Dispatch{
		prefix: []byte("The number of different routes from "),
		parser: parseQuest4,
	}

	errFormat = errors.New("I have no idea what you are talking about")
}

func ParseQuestion(line []byte) (Problem, error) {
	idx := bytes.Index(line, []byte{'.'})
	if idx == -1 {
		return nil, errFormat
	}

	id, err := strconv.Atoi(string(line[:idx]))
	if err != nil {
		return nil, errFormat
	}

	line = bytes.TrimSpace(line[idx+1:])

	for _, dispatch := range dispatchs {
		if bytes.HasPrefix(line, dispatch.prefix) {
			parse := dispatch.parser
			return parse(line[len(dispatch.prefix):], id)
		}
	}

	return nil, errFormat
}

func ParseGraph(input []byte) (*Graph, error) {
	if !bytes.HasPrefix(input, []byte("Graph:")) {
		return nil, errFormat
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
			return nil, errFormat
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

// example input: "C to C with a distance of less than 30."
func parseQuest4(line []byte, id int) (Problem, error) {
	ret := new(Quest4)
	ret.id = id

	if line[0] < 'A' || line[0] > 'Z' {
		return nil, errFormat
	}
	ret.from = line[0]
	line = line[1:]

	str := " to "
	if !bytes.HasPrefix(line, []byte(str)) {
		return nil, errFormat
	}
	line = line[len(str):]

	if line[0] < 'A' || line[0] > 'Z' {
		return nil, errFormat
	}
	ret.to = line[0]
	line = line[1:]

	str = " with a distance of less than "
	if !bytes.HasPrefix(line, []byte(str)) {
		return nil, errFormat
	}
	line = line[len(str):]

	idx := bytes.IndexByte(line, '.')
	if idx != len(line)-1 {
		return nil, errFormat
	}

	dist, err := strconv.Atoi(string(line[:idx]))
	if err != nil || dist < 0 {
		return nil, errFormat
	}

	ret.dist = dist
	return ret, err
}

// example input: "A to B."
func parseQuest3(line []byte, id int) (Problem, error) {
	ret := new(Quest3)
	ret.id = id

	str := "* to *."
	if len(line) != len(str) {
		return nil, errFormat
	}

	for i := 0; i < len(str); i++ {
		if str[i] != '*' && str[i] != line[i] {
			return nil, errors.New("I have no idea what you are talking about")
		}
	}

	ret.from = line[0]
	ret.to = line[5]
	return ret, nil
}

// example input: "C and ending at C with a maximum of 3 stops."
func parseQuest2(line []byte, id int) (Problem, error) {
	ret := new(Quest2)
	ret.id = id
	err := errFormat

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
	str = " stops."
	idx := bytes.Index(line, []byte(str))
	if idx == -1 {
		return nil, err
	}
	if idx+len(str) != len(line) {
		return nil, err
	}
	ret.stops, err = strconv.Atoi(string(line[:idx]))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// line format similar to A-B-C.
func parseQuest1(line []byte, id int) (Problem, error) {
	err := errFormat
	ret := new(Quest1)
	ret.id = id
	ret.input = make([]byte, 0, 5)
	for i := 0; i < len(line); {
		if line[i] < 'A' || line[i] > 'Z' {
			return nil, err
		}

		ret.input = append(ret.input, line[i])
		i++
		if i >= len(line) {
			return nil, err
		}

		if line[i] == '-' {
			i++
		} else if line[i] == '.' {
			break
		} else {
			return nil, err
		}
	}
	return ret, nil
}
