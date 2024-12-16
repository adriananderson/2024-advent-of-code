package main

import (
	"github.com/adriananderson/2024-advent-of-code/day16"
	"testing"
)

func Test16FirstSmall(t *testing.T) {
	if day16.Part1("day16/day16-test.txt") != 7036 {
		t.Error("maybe")
	}
}

func Test16FirstBig(t *testing.T) {
	if day16.Part1("day16/day16-test2.txt") != 11048 {
		t.Error("maybe")
	}
}

func Test16FSecondSmall(t *testing.T) {
	if day16.Part2("day16/day16-test.txt") != 45 {
		t.Error("maybe")
	}
}

func Test16SecondBig(t *testing.T) {
	if day16.Part2("day16/day16-test2.txt") != 64 {
		t.Error("maybe")
	}
}
