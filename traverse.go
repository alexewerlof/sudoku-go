package main

/*
 * All these traverse funcs are similar.
 * they get a board and a func.
 * they pass values to the func and
 * if the func returns something
 * truthy, it'll break the loop.
 */

type Traverser func(val, x, y int) (ret int, breakLoop bool)

func forRow(board Board, y int, fn Traverser) (int, bool) {
	return forRange(0, 9, 1, func(x int) {
		fn(board.getCell(x, y), x, y)
	})
}

func forCol(board Board, x int, fn Traverser) (int, bool) {
	return forRange(0, 9, 1, func(y int) {
		fn(board.getCell(x, y), x, y)
	})
}

func forRange(start, stop, step int, fn Traverser) (int, bool) {
	for i := start; i < stop; i = i + step {
		ret, breakLoop := fn(i)
		if breakLoop {
			return 0, ret
		}
	}
	return 0, false
}

func times(max int, fn Traverser) (int, bool) {
	return forRange(0, max, 1, fn)
}

func forRange2D(x1, x2, xDelta, y1, y2, yDelta int, fn Traverser) (int, bool) {
	return forRange(y1, y2, yDelta, func(y int) (int, bool) {
		return forRange(x1, x2, xDelta, func(x int) (int, bool) {
			return fn(x, y)
		})
	})
}

func forEveryCell(board Board, fn Traverser) (int, bool) {
	return forRange2D(0, 9, 1, 0, 9, 1, func(x, y int) (int, bool) {
		return fn(board.getCell(x, y), x, y)
	})
}

func forHouse(board Board, xLeft, yTop int, fn Traverser) (int, bool) {
	return forRange2D(xLeft, xLeft+3, 1, yTop, yTop+3, 1, func(x, y int) (int, bool) {
		return fn(board.getCell(x, y), x, y)
	})
}
