package solutions

func uniqueMorseRepresentations(words []string) int {
	morseCode := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	count := map[string]int{}

	for _, w := range words {
		count[concatMorse(w, &morseCode)]++
	}

	return len(count)
}

func concatMorse(words string, morseCode *[26]string) string {
	var rs string
	for _, c := range words {
		rs += morseCode[c-'a']
	}

	return rs
}
