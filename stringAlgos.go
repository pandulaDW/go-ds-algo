package main

import (
	"unicode"
)

// Best not to use len function on strings, as it returns a the number
// of elements in the byte slice. It would work for ASCII encoded text,
// but not on unicode characters that spans more than 1 byte

func lengthOfString(s string) int {
	length := 0
	for range s {
		length++
	}
	return length
}

// This would only work for ASCII characters
func makeUpperCase(s string) string {
	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'a' <= c && c <= 'z' {
			c = c - ('a' - 'A')
		}
		b = append(b, c)
	}
	return string(b)
}

// This would only work for ASCII characters
func makeLowerCase(s string) string {
	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c = c + ('a' - 'A')
		}
		b = append(b, c)
	}
	return string(b)
}

// make a map with each rune character in a given string
func makeMapFromString(s string) map[rune]int {
	m := map[rune]int{}

	for _, val := range s {
		m[val] = 0
	}

	return m
}

// Returns number of unique vowels and number of unique consants respectively
func countVowelAndConsonant(s string) (int, int) {
	vowels := makeMapFromString("aeiou")
	consonants := makeMapFromString("bcdfghjklmnpqrstvwxyz")

	for _, c := range s {
		// check if the character is an alphabet character
		if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' {

			// check if the character is a vowel
			switch {
			case c == 'A' || c == 'a':
				vowels['a']++
			case c == 'E' || c == 'e':
				vowels['e']++
			case c == 'I' || c == 'i':
				vowels['i']++
			case c == 'O' || c == 'o':
				vowels['o']++
			case c == 'U' || c == 'u':
				vowels['u']++
			default:
				consonants[unicode.ToLower(c)]++ // else, it should be consonant
			}
		}
	}

	vowelCount, consonantCount := 0, 0

	for _, val := range vowels {
		if val > 0 {
			vowelCount++
		}
	}

	for _, val := range consonants {
		if val > 0 {
			consonantCount++
		}
	}

	return vowelCount, consonantCount
}

// checks if a word contains non alphabet characters
func isValidWord(s string) bool {
	for _, val := range s {
		if !('A' <= val && val <= 'Z' || 'a' <= val && val <= 'z') {
			return false
		}
	}
	return true
}

// works for unicode as well
func reverseString(s string) string {
	runeSlice := make([]rune, 0, len(s))

	// create a rune slice from the underlying byte slice
	for _, val := range s {
		runeSlice = append(runeSlice, val)
	}

	reversed := make([]rune, len(runeSlice))

	for i := 0; i < len(runeSlice); i++ {
		reversed[i] = runeSlice[len(runeSlice)-i-1]
	}

	return string(reversed)
}

func compareString(s1, s2 string) int {
	slice1, slice2 := []rune(s1), []rune(s2)
	minLength := len(slice1)

	if len(slice1) > len(slice2) {
		minLength = len(slice2)
	}

	for i := 0; i < minLength; i++ {
		if slice1[i] < slice2[i] {
			return -1
		}

		if slice1[i] > slice2[i] {
			return 1
		}
	}

	// if the two strings are equal, then check the length
	if len(slice1) > len(slice2) {
		return 1
	}

	if len(slice1) < len(slice2) {
		return -1
	}

	// if lengths are also matching, then return 0
	return 0
}

// check if a word is a palindrome
func isPalindrome(s string) bool {
	word := []rune(s)

	for i, j := 0, len(word)-1; i < len(word); i, j = i+1, j-1 {
		if i >= j {
			break
		}
		if word[i] != word[j] {
			return false
		}
	}
	return true
}

// check if two given words are anagrams
func isAnagram(s1, s2 string) bool {
	word1 := []rune(s1)
	word2 := []rune(s2)
	elemMap := map[rune]int{}

	if len(word1) != len(word2) {
		return false
	}

	// iterate the first word and populate the map
	for i := 0; i < len(word1); i++ {
		c := word1[i]
		if _, ok := elemMap[c]; !ok {
			elemMap[c] = 1
		} else {
			elemMap[c]++
		}
	}

	// iterate the second word and modify the map
	for i := 0; i < len(word2); i++ {
		c := word2[i]
		if _, ok := elemMap[c]; ok {
			elemMap[c]--
		} else {
			return false // if a new letter is found return straight away
		}
	}

	// iterate the map and see if each character has only zero value
	for _, val := range elemMap {
		if val != 0 {
			return false
		}
	}
	return true
}

// finds the number of duplicate letters
func findDuplicateLetters(s string) []string {
	intSlice := make([]int, 0, len(s))

	// create a rune slice from the underlying byte slice
	for _, val := range s {
		intSlice = append(intSlice, int(val))
	}

	// get the duplicates
	duplicateInts := findDuplicateElements(intSlice)

	// convert the int slice to string slice
	duplicateLetters := make([]string, len(duplicateInts))
	for i, val := range duplicateInts {
		duplicateLetters[i] = string(val)
	}

	return duplicateLetters
}

// Only considering lowercase alphabet letters to reduce complexity
func findDuplicateLettersBitwise(s string) []string {
	var h, x int32

	result := make([]string, 0, len(s))

	for i := 0; i < len(s); i++ {
		offset := s[i] - 97
		x = 1
		x = x << offset
		if x&h > 0 {
			result = append(result, string(s[i]))
		} else {
			h = x | h
		}
	}

	return result
}

// finding permutations
/*
s := "ABC"
	word, length := []byte(s), len(s)
	results := make([]string, 0, factorial(length))

	// flag array to return the correct letter
	flag := make([]byte, length)

	// function to return the correct letter
	flagLetter := func() byte {
		for i := 0; i < length; i++ {
			if flag[i] == 0 {
				flag[i] = 1
				return word[i]
			}
		}
		return 0
	}

	// define the recursive function
	var traverse func(k int)

	// current permutation array
	permutation := make([]byte, 0, length)

	// recursive function implementation
	traverse = func(k int) {
		if k == length {
			results = append(results, string(permutation))
			return
		}
		permutation = append(permutation, flagLetter())
		traverse(k + 1)
	}

	traverse(0)

	fmt.Println(string(permutation))

}

func factorial(n int) int {
	if n < 2 {
		return 1
	}
	return n * factorial(n-1)
}
*/
