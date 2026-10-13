package main

func Ponc(c byte) bool {
	return c == '.' || c == ',' || c == '!' || c == '?' || c == ':' || c == ';'
}

func Ponctuation(text string) string {
	var resu []byte

	for i := 0; i < len(text); i++ {
		c := text[i]

		if Ponc(c) {
			for len(resu) > 0 && resu[len(resu)-1] == ' ' {
				resu = resu[:len(resu)-1]
			}

			resu = append(resu, c)

			if i+1 < len(text) && !Ponc(text[i+1]) && text[i+1] != ' ' {
				resu = append(resu, ' ')
			}
		} else {
			resu = append(resu, c)
		}
	}
	return string(resu)
}