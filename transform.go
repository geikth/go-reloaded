package main

import (
	"strconv"
	"strings"
)

func Transform(text string) string {
	mots := strings.Fields(text)
	var resultat []string

	for i := 0; i < len(mots); i++ {
		mot := mots[i]
		if mot == "(up)" {
			appliquer(resultat, 1, ToUpper)
		} else if mot == "(low)" {
			appliquer(resultat, 1, ToLower)
		} else if mot == "(cap)" {
			appliquer(resultat, 1, Cap)
		} else if mot == "(bin)" {
			appliquer(resultat, 1, Bin)
		}else if mot == "(hex)" {
			appliquer(resultat, 1, Hex)
		} else if (mot == "(up," || mot == "(low," || mot == "(cap,") && i+1 < len(mots) {
			n, err := strconv.Atoi(strings.TrimSuffix(mots[i+1], ")"))
			if err != nil {
				resultat = append(resultat, mot)
				continue
			}
			switch mot {
			case "(up,":
				appliquer(resultat, n, ToUpper)
			case "(low,":
				appliquer(resultat, n, ToLower)
			case "(cap,":
				appliquer(resultat, n, Cap)
			}
			i++
		} else {
			resultat = append(resultat, mot)
		}
	}
	Aan(resultat)
	return Ponctuation(strings.Join(resultat, " "))
}

func appliquer(liste []string, n int, f func(string) string) {
	debut := len(liste) - n
	if debut < 0 {
		debut = 0
	}
	for j := debut; j < len(liste); j++ {
		liste[j] = f(liste[j])
	}
}