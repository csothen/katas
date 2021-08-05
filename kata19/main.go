package kata19

import (
	"bufio"
	"fmt"
	"net/http"
)

const (
	dictionaryUrl = "http://codekata.com/data/wordlist.txt"
)

type Word struct {
	word    string
	visited bool
}

func Execute(start, end string) []string {
	dictionary, err := buildDictionary()
	if err != nil {
		panic(err)
	}
	return buildChain(start, end, dictionary)
}

func buildChain(start, end string, dictionary []Word) []string {
	chain := []string{start}
	current := start

	for {
		fmt.Println(current)
		word := findWord(current, end, dictionary)
		if word == nil {
			return []string{}
		}

		current = *word
		chain = append(chain, current)
		if current == end {
			break
		}
	}

	return chain
}

func findWord(start, end string, dictionary []Word) *string {
	left, right := 0, len(dictionary)-1
	auxDictionary := dictionary
	return binarySearch(start, end, auxDictionary, left, right)
}

func binarySearch(start, end string, dictionary []Word, left, right int) *string {
	if right < left {
		return nil
	}

	mid := left + (right-left)/2

	if isValid(dictionary[mid], start, end) {
		return &dictionary[mid].word
	}
	dictionary[mid].visited = true

	if dictionary[mid].word > end {
		return binarySearch(start, end, dictionary, left, mid-1)
	}

	// Otherwise it is to the right
	return binarySearch(start, end, dictionary, mid+1, right)
}

func isValid(w Word, current, target string) bool {
	if w.visited || len(w.word) != len(target) {
		return false
	}

	a := []rune(current)
	b := []rune(target)
	count := 0

	for i, c := range w.word {
		if a[i] != c && b[i] == c {
			count = count + 1
		}
	}

	return count == 1
}

func buildDictionary() ([]Word, error) {
	dictionary := []Word{}

	resp, err := http.Get(dictionaryUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		dictionary = append(dictionary, Word{word: scanner.Text(), visited: false})
	}

	// The dictionary is ordered by ASCII value
	return dictionary, nil
}
