package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfString(t *testing.T) {
	// testing only ascii characters
	s := "hello, world"
	assert.Equal(t, 12, lengthOfString(s))

	// testing with unicode characters
	s = "日本語ab!!"
	assert.Equal(t, 7, lengthOfString(s))
}

func TestMakeUpperCase(t *testing.T) {
	s := "Hello, World"
	assert.Equal(t, "HELLO, WORLD", makeUpperCase(s))
}

func TestMakeLowerCase(t *testing.T) {
	s := "Hello, World"
	assert.Equal(t, "hello, world", makeLowerCase(s))
}

func TestCountVowelAndConsonant(t *testing.T) {
	s := "Hello, World apple unfolds ice cream with me"
	vowels, consonants := countVowelAndConsonant(s)
	assert.Equal(t, 5, vowels)
	assert.Equal(t, 12, consonants)
}

func TestIsValidWord(t *testing.T) {
	s := "Pandula"
	assert.Equal(t, true, isValidWord(s))

	s = "pandula2"
	assert.Equal(t, false, isValidWord(s))

	s = "pandula!"
	assert.Equal(t, false, isValidWord(s))
}

func TestReverseString(t *testing.T) {
	// ASCII
	s := "Pandula"
	assert.Equal(t, "aludnaP", reverseString(s))

	// Unicode
	s = "hello日本"
	assert.Equal(t, "本日olleh", reverseString(s))
}

func TestFindDuplicateLetters(t *testing.T) {
	s := "helloo日本本"
	assert.ElementsMatch(t, []string{"o", "本", "l"}, findDuplicateLetters(s))
}

func TestFindDuplicateLettersBitwise(t *testing.T) {
	s := "finding"
	assert.ElementsMatch(t, []string{"i", "n"}, findDuplicateLettersBitwise(s))
}

func BenchmarkDuplicateMap(t *testing.B) {
	s := "finding"
	for i := 0; i < t.N; i++ {
		findDuplicateLetters(s)
	}
}

func BenchmarkDuplicateBitwise(t *testing.B) {
	s := "finding"
	for i := 0; i < t.N; i++ {
		findDuplicateLettersBitwise(s)
	}
}

func TestCompareStrings(t *testing.T) {
	s1, s2 := "hello", "lops"
	assert.Equal(t, -1, compareString(s1, s2))

	s1, s2 = "hello", "Hello"
	assert.Equal(t, 1, compareString(s1, s2))

	s1, s2 = "helloo", "hello"
	assert.Equal(t, 1, compareString(s1, s2))

	s1, s2 = "hello", "hello"
	assert.Equal(t, 0, compareString(s1, s2))

	s1, s2 = "hell", "hello"
	assert.Equal(t, -1, compareString(s1, s2))
}

func TestIsPalindrome(t *testing.T) {
	// odd case
	s := "rotator"
	assert.Equal(t, true, isPalindrome(s))

	// even case
	s = "anna"
	assert.Equal(t, true, isPalindrome(s))

	// incorrect case
	s = "haha"
	assert.Equal(t, false, isPalindrome(s))
}

func TestIsAnagram(t *testing.T) {
	// all unique characters
	s1 := "listen"
	s2 := "silent"
	assert.Equal(t, true, isAnagram(s1, s2))

	// includes duplicated characters
	s1 = "a gentleman"
	s2 = "elegant man"
	assert.Equal(t, true, isAnagram(s1, s2))

	// length wrong case
	s1 = "pandula"
	s2 = "wrong"
	assert.Equal(t, false, isAnagram(s1, s2))

	// length correct wrong case
	s1 = "hello world"
	s2 = "wrong babee"
	assert.Equal(t, false, isAnagram(s1, s2))
}
