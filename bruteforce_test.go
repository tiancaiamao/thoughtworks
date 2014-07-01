package main

import (
	"testing"
)

func TestBruteForceStop(t *testing.T) {
	g := makeGraph()
	within, _ := BruteForceNStops(g, 3, 'C', 'C')
	_, exactly := BruteForceNStops(g, 4, 'A', 'C')
	if within != 2 {
		t.Errorf("BruteForceNStop error, case (C C)")
	}
	if exactly != 3 {
		t.Errorf("BruteForceNStop error, case (A C)")
	}
}
