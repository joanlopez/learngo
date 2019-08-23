// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
		rubik     = puzzle{title: "rubik's cube", price: 5}
	)

	// thanks to the printer interface we can add different types of values
	// to the list.
	//
	// only rule: they need to implement the `printer` interface.
	// to do that: each type needs to have a print method.

	var store list
	store = append(store, &minecraft, &tetris, mobydick, rubik)
	store.print()
}
