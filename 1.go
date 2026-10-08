package main

func Capitalize(data string) string {
	b := []byte(data)
	inWord := false

	for i, r := range b {
		switch {
		case r >= 'a' && r <= 'z':
			if !inWord {
				b[i] = r - 32
				inWord = true
			} else {
				b[i] = r
			}
		case r >= 'A' && r <= 'Z':
			if !inWord {
				inWord = true
			} else {
				b[i] = r + 32
			}
		case r >= '0' && r <= '9':
			if !inWord {
				inWord = true
			}
		default:
			inWord = false
		}
	}
	return string(b)
}


func ToLower(s string) string {
	b := []byte(s)
	for i, r := range b {
		if r >= 'A' && r <= 'Z' {
			b[i] = r + 32
		}
	}
	return string(b)
}



func IsLower(s string) bool {
	for _, c := range s {
		if c < 'a' || c > 'z' {
			return false
		}
	}
	return true
}

func ToUpper(s string) string {
	b := []byte(s)
	for i, r := range b {
		if r >= 'a' && r <= 'z' {
			b[i] = r - 32
		}
	}
	return string(b)
}