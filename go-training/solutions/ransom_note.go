package solutions

func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteC := [26]byte{}
	count := len(ransomNote)

	for _, c := range ransomNote {
		ransomNoteC[c-'a']++
	}

	for _, c := range magazine {
		if ransomNoteC[c-'a'] > 0 {
			count--
			ransomNoteC[c-'a']--
		}
	}

	return count <= 0
}
