package main

func Lettre(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func Apo(text string) string {
	var resu []byte
	ouvrante := true
	//ouvrante vérifie si l'apostrophe est ouvrante ou fermante. elle commence forcément ouverante
	for i := 0; i < len(text); i++ {
		c := text[i]

		if c != '\'' {
			resu = append(resu, c)
			continue
		}

		if i > 0 && i+1 < len(text) && Lettre(text[i-1]) && Lettre(text[i+1]) {
			resu = append(resu, c)
			continue
		}

		if ouvrante {
			resu = append(resu, c)
			for i+1 < len(text) && text[i+1] == ' ' {
				i++
			}
		} else {
			for len(resu) > 0 && resu[len(resu)-1] == ' ' {
				resu = resu[:len(resu)-1]
			}
			resu = append(resu, c)
		}
		ouvrante = !ouvrante
	}
	return string(resu)
}
