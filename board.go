package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Board interface {
	getCell(x, y int) int
	setCell(x, y, val int)
}

type CellBoard struct {
	cells [9][9]int
}

type DiffBoard struct {
	parent Board
	x      int
	y      int
	val    int
}

func (b *CellBoard) getCell(x, y int) int {
	return b.cells[x][y]
}

func (b *CellBoard) setCell(x, y, val int) {
	b.cells[x][y] = val
}

func (b *DiffBoard) getCell(x, y int) int {
	if x == b.x && y == b.y {
		return b.val
	}
	return b.parent.getCell(x, y)
}

func (b *DiffBoard) setCell(x, y, val int) {
	if x == b.x && y == b.y {
		b.val = val
	} else {
		b.parent.setCell(x, y, val)
	}
}

func normalizeCell(val string) int {
	n, err := strconv.Atoi(val)
	if err != nil || n < 1 || n > 9 {
		return 0
	}
	return n
}

func colorDigit(digit int) string {
	return strconv.Itoa(digit)
}

/*
func colorDigit(digit int) {
	switch digit {
	case 1:
		return chalk.green(digit)
	case 2:
		return chalk.yellow(digit)
	case 3:
		return chalk.cyan(digit)
	case 4:
		return chalk.white(digit)
	case 5:
		return chalk.redBright(digit)
	case 6:
		return chalk.yellowBright(digit)
	case 7:
		return chalk.blueBright(digit)
	case 8:
		return chalk.magentaBright(digit)
	case 9:
		return chalk.cyanBright(digit)
	default:
		return chalk.red(digit)
	}
}
*/

// formerly boardToString
func String(board Board) string {
	var ret bytes.Buffer
	for y := 0; y < 9; y++ {
		if y != 0 && (y%3) == 0 {
			ret.WriteString("- - - + - - - + - - - \n")
		}
		for x := 0; x < 9; x++ {
			if x != 0 && (x%3) == 0 {
				ret.WriteString("| ")
			}
			ret.WriteString(colorDigit(board.getCell(x, y)) + " ")
		}
		ret.WriteString("\n")
	}
	return ret.String()
}

func FillBoard(board Board, rows string) (err error) {

	if len(rows) != 81 {
		return fmt.Errorf("Expecting exactly 81 chars found %d", len(rows))
	}

	var x, y, val int
	for r := 0; r < 81; r++ {
		x = r % 9
		y = int(float64(r) / 9)
		val = normalizeCell(string(rows[r]))
		board.setCell(x, y, val)
	}

	return nil
}
