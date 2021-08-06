package kata19

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	SEARCH_LIMIT = 15000
)

var (
	candidates = []string{}
	tracker    = make(map[string]interface{})
	count      = 0
)

var (
	ErrDifferentLengths    = fmt.Errorf("words must be of same length")
	ErrWordNotInDictionary = fmt.Errorf("word is not in the dictionary")
	ErrChainNotFound       = fmt.Errorf("couldn't find a valid chain")
	ErrSearchLimitReached  = fmt.Errorf("exceeded search limit")
)

func Execute(start, end string) []string {
	if len(start) != len(end) {
		log.Fatal(ErrDifferentLengths)
	}

	if start == end {
		return []string{start}
	}

	dictionary, err := buildDictionary(len(start))
	if err != nil {
		log.Fatal("Error building dictionary: ", err)
	}

	if err := checkExists(dictionary, start); err != nil {
		log.Fatal(err, start)
	}

	if err := checkExists(dictionary, end); err != nil {
		log.Fatal(err, end)
	}

	candidates = append([]string{}, start)
	tracker = make(map[string]interface{})
	tracker[start] = nil
	count = 0

	words, err := buildChain(end, dictionary)
	if err != nil {
		log.Fatal("Error building word chain: ", err)
	}

	return words
}

func checkExists(dictionary []string, word string) error {
	for _, w := range dictionary {
		if w == word {
			return nil
		}
	}

	return ErrWordNotInDictionary
}

func buildChain(end string, dictionary []string) ([]string, error) {
	for len(candidates) > 0 {
		count++
		if count > SEARCH_LIMIT {
			return nil, ErrSearchLimitReached
		}
		candidate := candidates[0]
		candidates = candidates[1:]

		for _, word := range dictionary {
			_, seen := tracker[word]
			if !seen && adjacent(candidate, word) {
				tracker[word] = candidate
				if end == word {
					return result(word), nil
				}
				candidates = append(candidates, word)
			}
		}
	}
	return nil, ErrChainNotFound
}

func adjacent(a, b string) bool {
	diff := 0
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff == 1
}

func result(word string) []string {
	result := []string{word}
	current := tracker[word]

	for current != nil {
		result = append([]string{current.(string)}, result...)
		current = tracker[current.(string)]
	}

	return result
}

func buildDictionary(length int) ([]string, error) {
	dictionary := []string{}

	f, err := os.Open("../wordlist.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == length {
			dictionary = append(dictionary, word)
		}
	}

	return dictionary, nil
}
