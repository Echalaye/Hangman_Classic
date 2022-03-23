package piscine

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func RecupToSommaire() string {

	dico, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("L'erreur est : %v", err.Error())
		os.Exit(1)
	}
	dictionnaire := string(dico)
	mots := strings.Split(dictionnaire, "\n")
	random := NombreAleatoire(len(mots))
	return mots[random]

}
