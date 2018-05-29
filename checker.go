package main

type Checker map[int]bool

func NewChecker() *Checker {
	ret := Checker{}
	return &ret
}

func (c Checker) Has(val int) bool {
	return c[val]
}

func (c Checker) Add(val int) {
	c[val] = true
}

func (c Checker) Reset() {
	for val := 1; val <= 9; val++ {
		c[val] = false
	}
}

func (c Checker) Full() bool {
	for val := 1; val <= 9; val++ {
		if !c.Has(val) {
			return false
		}
	}
	return true
}
