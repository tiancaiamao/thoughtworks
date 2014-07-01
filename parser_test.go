package main

import (
	"bytes"
	"testing"
)

func TestParseQuest1(t *testing.T) {
	prob, err := parseQuest1([]byte("A-B-C."), 0)
	if err != nil {
		t.Errorf("parseQuest1 not implement right")
	}
	q, ok := prob.(*Quest1)
	if !ok || !bytes.Equal([]byte("ABC"), q.input) {
		t.Errorf("parseQuest1 not implement right")
	}

	wrongTestCase := []string{
		"AB",
		"A-B",
		"a-b",
		"asdflksadjfe",
	}
	for i, testCase := range wrongTestCase {
		prob, err = parseQuest1([]byte(testCase), 0)
		if err == nil {
			t.Errorf("parseQuest1 cann't cope with invalid input for case %d\n", i)
		}
	}
}

func TestParseQuest2(t *testing.T) {
	prob, err := parseQuest2([]byte("C and ending at C with a maximum of 3 stops."), 0)
	if err != nil {
		t.Errorf("parseQuest2 not implement right")
	}

	q, ok := prob.(*Quest2)
	if !ok || q.start != 'C' || q.end != 'C' || q.exactly == true || q.stops != 3 {
		t.Errorf("parseQuest2 not implement right")
	}

	wrongTestCase := []string{
		"C and ending at C with a maximum of 3 stops.\n",
		"c and ending at C with a maximum of 3 stops.",
		"C and ending at C with a maximum of 3",
		"C and ending at C with maximum of 7-",
		"asdflksadjfe",
	}
	for i, testCase := range wrongTestCase {
		prob, err = parseQuest2([]byte(testCase), 0)
		if err == nil {
			t.Errorf("parseQuest2 cann't cope with invalid input for case %d\n", i)
		}
	}
}

func TestParseQuest4(t *testing.T) {
	prob, err := parseQuest4([]byte("C to C with a distance of less than 30."), 0)
	if err != nil {
		t.Errorf("parseQuest4 not implement right")
	}

	q, ok := prob.(*Quest4)
	if !ok || q.from != 'C' || q.to != 'C' || q.dist != 30 {
		t.Errorf("parseQuest4 not implement right")
	}

	wrongTestCase := []string{
		"C to C with a distance of less than 30.\n",
		"B to C with a distance of less than -1.",
		"C to C with	 distance of than 3.",
		"asdflksadjfe",
	}
	for i, testCase := range wrongTestCase {
		prob, err = parseQuest4([]byte(testCase), 0)
		if err == nil {
			t.Errorf("parseQuest4 cann't cope with invalid input for case %d\n", i)
		}
	}
}
