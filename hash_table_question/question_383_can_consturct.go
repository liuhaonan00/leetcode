package hash_table_question

func CanConstruct(ransomNote string, magazine string) bool {
	return canConstruct(ransomNote, magazine)
}

// 主要使用hash table 计数，遍历magazine 里边的每一个char

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) == 0 {
		return true
	}

	if len(magazine) == 0 || len(ransomNote) > len(magazine) {
		return false
	}

	magazineDict := map[rune]int{}

	for _, r := range magazine {
		if _, ok := magazineDict[r]; ok {
			magazineDict[r]++
		} else {
			magazineDict[r] = 1
		}

	}

	for _, r := range ransomNote {
		if magazineDict[r] <= 0 {
			return false
		}
		magazineDict[r]--
	}
	return true
}
