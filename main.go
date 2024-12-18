package main

import (
	"fmt"
	"github.com/adriananderson/2024-advent-of-code/day01"
	"github.com/adriananderson/2024-advent-of-code/day02"
	"github.com/adriananderson/2024-advent-of-code/day03"
	"github.com/adriananderson/2024-advent-of-code/day04"
	"github.com/adriananderson/2024-advent-of-code/day05"
	"github.com/adriananderson/2024-advent-of-code/day06"
	"github.com/adriananderson/2024-advent-of-code/day07"
	"github.com/adriananderson/2024-advent-of-code/day08"
	"github.com/adriananderson/2024-advent-of-code/day09"
	"github.com/adriananderson/2024-advent-of-code/day10"
	"github.com/adriananderson/2024-advent-of-code/day11"
	"github.com/adriananderson/2024-advent-of-code/day12"
	"github.com/adriananderson/2024-advent-of-code/day13"
	"github.com/adriananderson/2024-advent-of-code/day14"
	"github.com/adriananderson/2024-advent-of-code/day15"
	"github.com/adriananderson/2024-advent-of-code/day16"
	"github.com/adriananderson/2024-advent-of-code/day17"
	"github.com/adriananderson/2024-advent-of-code/day18"
	"time"
)

func main() {
	start := time.Now()

	//// Day 01
	fmt.Printf("Final result Day 01 part 1: %d\n", day01.Part1())
	fmt.Printf("Final result Day 01 part 2: %d\n", day01.Part2())
	//// Day 02
	fmt.Printf("Final result Day 02 part 1: %d\n", day02.Part1())
	fmt.Printf("Final result Day 02 part 2: %d\n", day02.Part2())
	//// Day 03
	fmt.Printf("Final result Day 03 part 1: %d\n", day03.Part1())
	fmt.Printf("Final result Day 03 part 2: %d\n", day03.Part2())
	//// Day 04
	fmt.Printf("Final result Day 04 part 1: %d\n", day04.Part1())
	fmt.Printf("Final result Day 04 part 2: %d\n", day04.Part2())
	//// Day 05
	fmt.Printf("Final result Day 05 part 1: %d\n", day05.Part1())
	fmt.Printf("Final result Day 05 part 2: %d\n", day05.Part2())
	//// Day 06
	fmt.Printf("Final result Day 06 part 1: %d\n", day06.Part1())
	fmt.Printf("Final result Day 06 part 2: %d\n", day06.Part2())
	//// Day 07
	fmt.Printf("Final result Day 07 part 1: %d\n", day07.Part1())
	fmt.Printf("Final result Day 07 part 2: %d\n", day07.Part2())
	//// Day 08
	fmt.Printf("Final result Day 08 part 1: %d\n", day08.Part1())
	fmt.Printf("Final result Day 08 part 2: %d\n", day08.Part2())
	//// Day 09
	fmt.Printf("Final result Day 09 part 1: %d\n", day09.Part1())
	fmt.Printf("Final result Day 09 part 2: %d\n", day09.Part2())
	//// Day 10
	fmt.Printf("Final result Day 10 part 1: %d\n", day10.Part1())
	fmt.Printf("Final result Day 10 part 2: %d\n", day10.Part2())
	//// Day 11
	fmt.Printf("Final result Day 11 part 1: %d\n", day11.Part1())
	fmt.Printf("Final result Day 11 part 2: %d\n", day11.Part2())
	//// Day 12
	fmt.Printf("Final result Day 12 part 1: %d\n", day12.Part1())
	fmt.Printf("Final result Day 12 part 2: %d\n", day12.Part2())
	//// Day 13
	fmt.Printf("Final result Day 13 part 1: %d\n", day13.Part1())
	fmt.Printf("Final result Day 13 part 2: %d\n", day13.Part2())
	//// Day 14
	fmt.Printf("Final result Day 14 part 1: %d\n", day14.Part1())
	fmt.Printf("Final result Day 14 part 2: %d\n", day14.Part2())
	//// Day 15
	fmt.Printf("Final result Day 15 part 1: %d\n", day15.Part1())
	fmt.Printf("Final result Day 15 part 2: %d\n", day15.Part2())
	//// Day 16
	fmt.Printf("Final result Day 16 part 1: %d\n", day16.Part1("day16/day16.txt"))
	fmt.Printf("Final result Day 16 part 2: %d\n", day16.Part2("day16/day16.txt"))
	//// Day 17
	fmt.Printf("Final result Day 17 part 1: %v\n", day17.Part1("day17/day17.txt"))
	fmt.Printf("Final result Day 17 part 2: %d\n", day17.Part2("day17/day17.txt"))
	//// Day 18
	fmt.Printf("Final result Day 18 part 1: %d\n", day18.Part1("day18/day18.txt", 70, 1024))
	fmt.Printf("Final result Day 18 part 2: %s\n", day18.Part2("day18/day18.txt", 70, 1024))

	fmt.Printf("... took %v\n", time.Since(start))
}
