package main

func Aan(liste []string) {
	for i := 0; i < len(liste)-1; i++ {
		if liste[i] == "a" || liste[i] == "A" {
			suivant := ToLower(liste[i+1])
			if suivant != "" && voyelle(suivant[0]) {
				if liste[i] == "a" {
					liste[i] = "an"
				} else {
					liste[i] = "An"
				}
			}
		}
	}
}

func voyelle(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'h'
}