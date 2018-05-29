package main

import "fmt"

func findFirstEmptyCell(board Board) (x, y int, found bool) {
	for x = 0; x < 9; x++ {
		for y = 0; y < 9; y++ {
			if board.getCell(x, y) == 0 {
				found = true
				return
			}
		}
	}
	return
}

func rowsFail(board Board) string {
	checker := NewChecker()
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			val := board.getCell(x, y)
			if val != 0 {
				if checker.Has(val) {
					return fmt.Sprintf("Row %d has at least two %ds", y, val)
				}
				checker.Add(val)
			}
		}

		checker.Reset()
	}
	return ""
}

func colsFail(board Board) string {
	checker := NewChecker()
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			val := board.getCell(x, y)
			if val != 0 {
				if checker.Has(val) {
					return fmt.Sprintf("Col %d has at least two %ds", x, val)
				}
				checker.Add(val)
			}
		}

		checker.Reset()
	}
	return ""
}

func housesFail(board Board) string {
	checker := NewChecker()
	for xLeft := 0; xLeft < 9; xLeft += 3 {
		for yTop := 0; yTop < 9; yTop += 3 {

			for x := xLeft; x < xLeft+3; x++ {
				for y := yTop; y < yTop+3; y++ {
					val := board.getCell(x, y)
					if val != 0 {

						if checker.Has(val) {
							return fmt.Sprintf("House at top %d and left %d has at least two %ds", xLeft, yTop, val)
						}

						checker.Add(val)
					}
				}
			}

			checker.Reset()
		}
	}
	return ""
}

var failCounter int
var winCounter int

func anyFail(board Board) string {
	fail := rowsFail(board)
	if fail != "" {
		return fail
	}
	fail = colsFail(board)
	if fail != "" {
		return fail
	}
	fail = housesFail(board)
	if fail != "" {
		return fail
	}
	return ""
}

func Solve(board Board) {
	fail := anyFail(board)
	if fail != "" {
		// fmt.Println(fail);
		failCounter++
		return
	}
	x, y, foundEmpty := findFirstEmptyCell(board)
	if foundEmpty {
		// fmt.Println(`Found an empty cell at ${x},${y}`);
		for val := 1; val <= 10; val++ {
			newBoard := DiffBoard{parent: board, x: x, y: y, val: val}
			Solve(&newBoard)
		}
	} else {
		winCounter++
		fmt.Println("Try", failCounter, "Win", winCounter)
		fmt.Println(board)
	}
}
