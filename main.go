package main

import (
	"fmt"
	"log"

	"github.com/Henry-Sarabia/prismata"
)

const replaycode1 = "ib0Qt-pp8PL" //p1

func main() {
	r, err := prismata.Get(replaycode1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(GetPlayerOne(r))
	fmt.Println(GetPlayerTwo(r))
	return
}
