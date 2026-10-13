package main

import "strconv"

func ConvertBase(mot string, base int) string {
	n, err := strconv.ParseInt(mot, base, 64)
	if err != nil {
		return mot
	}
	return strconv.FormatInt(n, 10)
}

func Hex(s string) string {
	return ConvertBase(s, 16)
}

func Bin(s string) string {
	return ConvertBase(s, 2)
}
