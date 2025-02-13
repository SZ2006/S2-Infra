package main

import (
	"fmt"
	"time"
)

func main() {
	uur := time.Now().Hour()
	if uur >= 7 && uur < 12 {
		fmt.Println("Goedemorgen! Welkom bij Fonteyn Vakantieparken")
	} else if uur >= 12 && uur < 18 {
		fmt.Println("Goedemiddag! Welkom bij Fonteyn Vakantieparken")
	} else if uur >= 18 && uur < 23 {
		fmt.Println("Goedenavond! Welkom bij Fonteyn Vakantieparken")
	} else {
		fmt.Println("Sorry, de parkeerplaats is ’s nachts gesloten")
	}
}
