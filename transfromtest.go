package main

import (
	"os"
	"strings"
	"testing"
)

func TestTransform(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// --- (hex) et (bin) ---
		{"hex", "1E (hex) files were added", "30 files were added"},
		{"bin", "It has been 10 (bin) years", "It has been 2 years"},
		{"hex invalide", "zzz (hex)", "zzz"},
		{"hex et bin ensemble", "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},

		// --- (up), (low), (cap) ---
		{"up", "Ready, set, go (up) !", "Ready, set, GO!"},
		{"low", "I should stop SHOUTING (low)", "I should stop shouting"},
		{"cap", "Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"cap sur un mot en majuscules", "HELLO (cap)", "Hello"},
		{"commande en début de texte", "(up) hello", "hello"},

		// --- (up, n), (low, n), (cap, n) ---
		{"up avec nombre", "This is so exciting (up, 2)", "This is SO EXCITING"},
		{"low avec nombre", "I AM VERY LOUD (low, 2) today", "I AM very loud today"},
		{"cap avec nombre", "it was the age of wisdom (cap, 3) ok", "it was the Age Of Wisdom ok"},
		{"nombre plus grand que le nombre de mots", "hello world (up, 5)", "HELLO WORLD"},
		{"nombre invalide", "hello (up, abc)", "hello (up, abc)"},

		// --- a -> an ---
		{"a devant une voyelle", "There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"a devant un h", "It was a honest mistake", "It was an honest mistake"},
		{"a devant une consonne", "I bought a car", "I bought a car"},
		{"a tout seul", "a", "a"},

		// --- ponctuation ---
		{"ponctuation collée au mot précédent", "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"points de suspension", "I was thinking ... You were right", "I was thinking... You were right"},
		{"groupe !?", "Really !?", "Really!?"},
		{"phrase du sujet", "Punctuation tests are ... kinda boring ,what do you think ?", "Punctuation tests are... kinda boring, what do you think?"},

		// --- apostrophes ---
		{"apostrophes, un mot", "I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"apostrophes, plusieurs mots", "As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"apostrophe dans un mot", "I don't know ' maybe '", "I don't know 'maybe'"},

		// --- cas généraux ---
		{"texte vide", "", ""},
		{"espaces multiples", "hello     world", "hello world"},
		{
			"tout ensemble",
			"this is so exciting (up, 2) ,and 1E (hex) files were added . I SHOUT (low) now ! there was a amazing rock (cap) ... I am exactly how they describe me: ' awesome ' and 10 (bin) years passed",
			"this is SO EXCITING, and 30 files were added. I shout now! there was an amazing Rock... I am exactly how they describe me: 'awesome' and 2 years passed",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Transform(tc.input)
			if got != tc.want {
				t.Errorf("Transform(%q)\n got: %q\nwant: %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestConvertBase(t *testing.T) {
	tests := []struct {
		mot  string
		base int
		want string
	}{
		{"1E", 16, "30"},
		{"1e", 16, "30"},
		{"10", 2, "2"},
		{"zzz", 16, "zzz"},
		{"102", 2, "102"},
	}

	for _, tc := range tests {
		got := ConvertBase(tc.mot, tc.base)
		if got != tc.want {
			t.Errorf("ConvertBase(%q, %d) = %q, want %q", tc.mot, tc.base, got, tc.want)
		}
	}
}

// TestSample lance Transform sur sample.txt et compare avec expected.txt.
func TestSample(t *testing.T) {
	entree, err := os.ReadFile("sample.txt")
	if err != nil {
		t.Fatalf("lecture de sample.txt impossible : %v", err)
	}
	attendu, err := os.ReadFile("expected.txt")
	if err != nil {
		t.Fatalf("lecture de expected.txt impossible : %v", err)
	}

	got := strings.TrimSpace(Transform(string(entree)))
	want := strings.TrimSpace(string(attendu))

	if got != want {
		t.Errorf("sample.txt\n got: %q\nwant: %q", got, want)
	}
}