package main


import (
	"fmt"
	"unicode"
)

const  exclamation = '!'
const  questionMark = '?'
const  hypen = '-'
const  apostrophe = '\''
const  point = '.'

func wordsCounter(sentence string) map[string]int {
	wordCounter := make(map[string]int)
	word := make([]rune, 0, 10)

	exclamationRune := rune(exclamation)
	questionMarkRune := rune(questionMark)
	hypenRune := rune(hypen)
        apostropheRune := rune(apostrophe)
	pointRune := rune(point)

	for i,v := range sentence {
		if unicode.IsLetter(v) || (len(word) > 0 && (v == hypenRune || v == apostropheRune)) {
			word = append(word, unicode.ToLower(v))
			if len(sentence) - 1 == i {
                                wordCounter[string(word)] += 1
			}
		} else if unicode.IsSpace(v) || exclamationRune == v || questionMarkRune == v || pointRune == v{
			if len(word) > 0 {
				wordCounter[string(word)] += 1
				word = make([]rune, 0, 10)
			}
		}
	}

	return wordCounter
}

func main() {
	input1 := "Strawberry short cake? Yum!"
	fmt.Printf("%v\n", wordsCounter(input1))

        input2 := "Dessert - mille-feuille cake"
	fmt.Printf("%v\n", wordsCounter(input2))

	input3 := "Mmm...mmm...decisions...decisions"
        fmt.Printf("%v\n", wordsCounter(input3))

	input4 := "Allie's Bakery: Sasha's Cakes"
	fmt.Printf("%v\n", wordsCounter(input4))
}
