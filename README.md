thoughtworks的编程测试，题目是trains，挂了，这是我提交的代码。

# build

first, [you need a golang envirenment](http://golang.org)

then, copy trains dir to $GOPATH

	go install trains

no extra dependency required, the executable will be in $GOPATH/bin

# run

	cd $GOPATH/src/trains
	../../bin/trains graphFile questionFile

if $GOPATH/bin is in your $PATH, that also work, of course 

	trains graphFile questionFile

Personally, I suggest you do that, Go is great!

you can slightly change graphFile and questionFile to do some test, but should remain the right format~

# other

to run unit test, just type:

	go test -v trains

to see I'm a honest interviewee, :-)

	git log

code struct:

	.
	├── bruteforce.go	
	├── bruteforce_test.go
	├── dijkstra.go		// implement dijkstra algorithm
	├── graph.go		// data struct represent of Graph in adjacency
	├── graphFile		// test file
	├── graph_test.go	
	├── main.go			// main file
	├── parser.go		// parse input into internal data struct
	├── questionFile	// test file
	└── solver.go		// solve the abstract problem
