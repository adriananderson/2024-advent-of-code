package main

import (
	"github.com/adriananderson/2024-advent-of-code/day07"
	"github.com/adriananderson/2024-advent-of-code/day16"
	"github.com/adriananderson/2024-advent-of-code/day17"
	"github.com/adriananderson/2024-advent-of-code/day18"
	"github.com/adriananderson/2024-advent-of-code/day19"
	"testing"
)

func Test7Second(t *testing.T) {
	val := day07.Part2()
	expected := 340362529351427
	if val != expected {
		t.Errorf("expected %d, got %d", expected, val)
	}
}

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

func Test17One(t *testing.T) {
	val := day17.Part1("day17/day17-test.txt")
	expected := "4,6,3,5,6,3,5,2,1,0"
	if val != expected {
		t.Errorf("expected %s, got %s", expected, val)
	}
}

func Test17Two(t *testing.T) {
	val := day17.Part2("day17/day17-test2.txt")
	expected := 117440
	if val != expected {
		t.Errorf("expected %d, got %d", expected, val)
	}
}

func Test17TwoReal(t *testing.T) {
	val := day17.Part2("day17/day17.txt")
	expected := 37222273957364
	if val != expected {
		t.Errorf("expected %d, got %d", expected, val)
	}
}

func Test18One(t *testing.T) {
	val := day18.Part1("day18/day18-test.txt", 6, 12)
	expected := 22
	if val != expected {
		t.Errorf("expected %d, got %d", expected, val)
	}
}

func Test18OneReal(t *testing.T) {
	val := day18.Part1("day18/day18.txt", 70, 1024)
	expected := 384
	if val != expected {
		t.Errorf("expected %d, got %d", expected, val)
	}
}

func Test18Two(t *testing.T) {
	val := day18.Part2("day18/day18-test.txt", 6, 12)
	expected := "6,1"
	if val != expected {
		t.Errorf("expected %s, got %s", expected, val)
	}
}

func Test18TwoReal(t *testing.T) {
	val := day18.Part2("day18/day18.txt", 70, 1024)
	expected := "36,10"
	if val != expected {
		t.Errorf("expected %s, got %s", expected, val)
	}
}

func Test19One(t *testing.T) {
	val := day19.Part1("day19/day19-test.txt")
	expected := 6
	if val != expected {
		t.Errorf("expected %d, got %d", expected, val)
	}
}

func Test19OneReal(t *testing.T) {
	val := day19.Part1("day19/day19.txt")
	expected := 311
	if val != expected {
		t.Errorf("expected %d, got %d", expected, val)
	}
}

func Test19Two(t *testing.T) {
	val := day19.Part2("day19/day19-test.txt")
	expected := 16
	if val != expected {
		t.Errorf("expected %d, got %d", expected, val)
	}
}

func Test19TwoReal(t *testing.T) {
	val := day19.Part2("day19/day19.txt")
	expected := 616234236468263
	if val != expected {
		t.Errorf("expected %d, got %d", expected, val)
	}
}
