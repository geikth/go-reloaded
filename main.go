package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	result := Transform(string(data))
	err = os.WriteFile(os.Args[2], []byte(result), 0644)
}

func Transform(text string) string {
	mots := strings.Fields(text)
	var resultat []string

	for i := 0; i < len(mots); i++ {
		if mots[i] == "(up)" {
			if len(resultat) > 0 {
				dernier := len(resultat) - 1
				resultat[dernier] = ToUpper(resultat[dernier])
			}
		} else {
			resultat = append(resultat, mots[i])
		}
	}
	return strings.Join(resultat, " ")
}