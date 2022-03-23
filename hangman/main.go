package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"piscine"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Vous n'avez pas entré de fichier correct.")
		return
	}
	pendu, err := ioutil.ReadFile("hangman/hangman.txt")
	if err != nil {
		fmt.Printf("L'erreur est : %v", err.Error())
		return
	}
	dessin := string(pendu)
	rand.Seed(time.Now().UnixNano())
	mot_cherche := (piscine.RecupToSommaire())
	tentatives := 10
	vrai := 0
	signe := 0
	compteur := 0
	save := []int{}
	tab_mot := []string{}
	tab_lettre := []string{}
	jeu := true
	pas_bon := false
	lettre_val := true
	mot_affichee := piscine.Affich(mot_cherche)
	nb_lettre_trouve := (len(mot_cherche) - 1) - (len(mot_cherche) / 2) + 1
	if len(mot_cherche) == 4 {
		nb_lettre_trouve = 3
	}
	var lettre string
	for jeu {
		vrai = 0
		lettre = ""
		fmt.Print("\nVous avez ", tentatives, " tentatives\n")
		fmt.Printf("\n")
		fmt.Print("Entrez un mot ou une lettre :")
		fmt.Scanln(&lettre)
		verif := []rune(lettre)
		for i := 0; i < len(verif); i++ {
			if rune(verif[i]) < 65 || rune(verif[i]) > 90 && rune(verif[i]) < 97 || rune(verif[i]) > 122 {
				fmt.Print("Tu n'as ni entré de lettre ni de mot.")
				pas_bon = true
				break
			} else {
				continue
			}
		}
		if !pas_bon {
			if len(lettre) > 1 {
				for m := 0; m < len(lettre); m++ {
					for i := 0; i < len(tab_lettre); i++ {
						tab_deja_lettre := []rune(lettre)
						if rune(tab_deja_lettre[m]) > 64 && rune(tab_deja_lettre[m]) < 91 {
							tab_deja_lettre[m] = rune(tab_deja_lettre[m] + 32)
						}
						if string(tab_lettre[i]) == string(lettre) {
							lettre_val = false
						}
					}
				}
			} else if len(lettre) == 1 {
				for n := 0; n < len(tab_lettre); n++ {
					tab_deja_lettre := []rune(lettre)
					if rune(tab_deja_lettre[0]) > 64 && rune(tab_deja_lettre[0]) < 91 {
						tab_deja_lettre[0] = rune(tab_deja_lettre[0] + 32)
					}
					if string(tab_lettre[n]) == string(lettre) {
						lettre_val = false
					}
				}
			}
			if len(lettre) == 0 {
				if compteur == 4 {
					tentatives = 0
					piscine.AffichePendu(dessin, tentatives)
					fmt.Print("\n")
					fmt.Println("Frère t'abuse !")
					jeu = false
				} else {
					fmt.Print("Vous n'avez rien écrit.")
					compteur += 1
				}
			} else if lettre_val {
				compteur = 0
				signe = 0
				if len(lettre) > 1 {
					for k := 0; k < len(verif); k++ {
						if rune(verif[k]) > 64 && rune(verif[k]) < 91 {
							tab_mot = append(tab_mot, string(rune(verif[k])+32))
						} else {
							tab_mot = append(tab_mot, string(verif[k]))
						}
					}
					lettre = ""
					for h := 0; h < len(tab_mot); h++ {
						lettre += tab_mot[h]
					}
					if lettre == string(mot_cherche[:len(mot_cherche)-1]) {
						jeu = false
						fmt.Println("Bravo tu as gagné !")
						fmt.Println(mot_cherche)
					} else {
						if tentatives == 2 {
							tentatives -= 2
							piscine.AffichePendu(dessin, tentatives)
							fmt.Print("\n")
							fmt.Println("Tu as perdu !")
							fmt.Println("Le mot était :", mot_cherche)
							jeu = false
						} else if tentatives == 1 {
							tentatives -= 1
							piscine.AffichePendu(dessin, tentatives)
							fmt.Print("\n")
							fmt.Println("Tu as perdu !")
							fmt.Println("Le mot était :", mot_cherche)
							jeu = false
						} else {
							tentatives -= 2
							piscine.AffichePendu(dessin, tentatives)
							fmt.Print("\n")
							fmt.Println("Ce n'est pas le bon mot. Il vous reste,", tentatives, "tentatives")
							piscine.Affichmotfaux(mot_affichee)
						}
					}
					tab_mot = []string{}
				} else {
					for i := 0; i < len(mot_cherche); i++ {
						if rune(verif[0]) > 64 && rune(verif[0]) < 91 {
							lettre = string(rune(verif[0]) + 32)
							if lettre == string(mot_cherche[i]) && lettre != string(mot_affichee[i]) {
								vrai = 1
								save = append(save, i)
								nb_lettre_trouve -= 1
							}
						} else if lettre == string(mot_cherche[i]) && lettre != string(mot_affichee[i]) {
							vrai = 1
							save = append(save, i)
							nb_lettre_trouve -= 1
						}
					}
					if tentatives == 1 && vrai == 0 {
						tentatives -= 1
						piscine.AffichePendu(dessin, tentatives)
						fmt.Print("\n")
						fmt.Println("Tu as perdu !")
						fmt.Println("Le mot était :", mot_cherche)
						jeu = false
					} else if nb_lettre_trouve == 0 {
						jeu = false
						fmt.Println("Bravo, tu as réussi !")
						fmt.Println("Le mot était bien :", mot_cherche)
					} else if vrai == 1 {
						piscine.AffichePendu(dessin, tentatives)
						mot_affichee = piscine.AffichIfVrai(lettre, mot_affichee, save)
						fmt.Print("\n")
					} else if vrai == 0 {
						tentatives--
						piscine.AffichePendu(dessin, tentatives)
						fmt.Print("\n")
						fmt.Println("Cette lettre n'est pas dans le mot. Il vous reste,", tentatives, "tentatives")
						piscine.Affichmotfaux(mot_affichee)
					} else {
						piscine.AffichePendu(dessin, tentatives)
						fmt.Println(mot_affichee)
						fmt.Print("\n")
					}
					save = []int{}
				}
				tab_lettre = append(tab_lettre, lettre)
			} else {
				fmt.Println("Vous avez déjà entré cette lettre.")
				piscine.AffichePendu(dessin, tentatives)
				lettre_val = true
			}
		} else {
			pas_bon = false
			signe += 1
			if signe == 2 {
				tentatives = 0
				piscine.AffichePendu(dessin, tentatives)
				fmt.Print("\n")
				fmt.Println("Tu as perdu !")
				fmt.Print("Fallait pas jouer au con !")
				jeu = false
			}
		}
	}
}
