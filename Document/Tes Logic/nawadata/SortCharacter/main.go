// Sort Character (50 Poin) (NDL010)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input one line of words (S) : ")
	scanner.Scan()

	vowel, consonant := convertCharacter(scanner.Text())

	fmt.Println("Vowel Characters :")
	fmt.Println(vowel)

	fmt.Println()
	fmt.Println("Consonant Characters :")
	fmt.Println(consonant)
}

func convertCharacter(word string) (string, string) {
	trimWord := strings.ReplaceAll(word, " ", "")
	var resultVowel []string
	var resultConsonant []string
	for i := 0; i < len(trimWord); i++ {
		str := string(trimWord[i])
		str = strings.ToLower(str)
		switch {
		case str == "a":
			resultVowel = append(resultVowel, str)
		case str == "i":
			resultVowel = append(resultVowel, str)
		case str == "u":
			resultVowel = append(resultVowel, str)
		case str == "e":
			resultVowel = append(resultVowel, str)
		case str == "o":
			resultVowel = append(resultVowel, str)
		default:
			resultConsonant = append(resultConsonant, str)
		}
	}
	vowel := strings.Join(resultVowel, "")
	consonant := strings.Join(resultConsonant, "")
	return vowel, consonant
}
