package main

import "contract/structs"

func main() {
	var (
		mobydick = structs.Book{Title: "Moby Dick", Price: 18.24}
		megaman  = structs.Game{Title: "Mega Man 3", Price: 33.99}
		metroid  = structs.Game{Title: "Metroid Prime", Price: 59.99}
    rubiks = structs.Puzzle{Title: "Metroid Prime", Price: 59.99, Difficulty: structs.Easy}
	)

	stock := structs.Stock{mobydick, megaman, metroid, rubiks}
	stock.Print()
}
