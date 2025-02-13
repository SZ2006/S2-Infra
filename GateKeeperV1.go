package main

import (
	"fmt"  
	"time" 
)

func toonBericht() {
	toegestaneKentekens := map[string]bool{
		"AB-123-CD": true,
		"EF-456-GH": true,
		"IJ-789-KL": true,
	}

	var kenteken string
	fmt.Print("Voer uw kenteken in: ")
	fmt.Scanln(&kenteken)

	if toegestaneKentekens[kenteken] {
		fmt.Println("Welkom op het parkeerterrein")
	} else {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
	}

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

func main() {
	toonBericht()
}
