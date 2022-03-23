package piscine

import (
	"fmt"
)

func AffichIfVrai(la_lettre string, motaafficher string, place []int) string {
	result := ""
	vrai := true
	for i := 0; i < len(motaafficher); i++ {
		vrai = true
		for j := 0; j < len(place); j++ {
			if i == place[j] {
				fmt.Printf("%v", string(la_lettre))
				result += string(la_lettre)
				vrai = false
			}
		}
		if vrai {
			fmt.Print(string(motaafficher[i]))
			result += string(motaafficher[i])
			vrai = false
		}
		fmt.Printf(" ")
	}
	return result
}
