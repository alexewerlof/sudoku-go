// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"log"
)

func main() {
	b := CellBoard{}
	err := FillBoard(&b, "956      "+
		"         "+
		" 8  463 7"+
		"6  28    "+
		"5  7  48 "+
		"2        "+
		"  5 7  4 "+
		"     372 "+
		"      6  ")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(String(&b))
	// b.solve()
}
