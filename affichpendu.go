package piscine

import (
	"fmt"
)

func AffichePendu(dessin string, numero int) {
	if numero == 9 {
		for i := 0; i < 77; i++ {
			fmt.Print(string(dessin[i]))
		}
	}
	if numero == 8 {
		for i := 78; i < 156; i++ {
			fmt.Print(string(dessin[i]))
		}
	}
	if numero == 7 {
		for i := 156; i < 235; i++ {
			fmt.Print(string(dessin[i]))
		}
	}
	if numero == 6 {
		for i := 235; i < 314; i++ {
			fmt.Print(string(dessin[i]))
		}
	}
	if numero == 5 {
		for i := 314; i < 394; i++ {
			fmt.Print(string(dessin[i]))
		}
	}
	if numero == 4 {
		for i := 394; i < 472; i++ {
			fmt.Print(string(dessin[i]))
		}
	}
	if numero == 3 {
		for i := 472; i < 551; i++ {
			fmt.Print(string(dessin[i]))
		}
	}
	if numero == 2 {
		for i := 551; i < 630; i++ {
			fmt.Print(string(dessin[i]))
		}
	}
	if numero == 1 {
		for i := 630; i < 709; i++ {
			fmt.Print(string(dessin[i]))
		}
	}
	if numero == 0 {
		for i := 709; i < 788; i++ {
			fmt.Print(string(dessin[i]))
		}
	}
}
