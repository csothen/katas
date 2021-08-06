package kata19

import (
	"bufio"
	"fmt"
	"os"
)

const (
	SEARCH_LIMIT = 20000 // Limit applied on the searches done
)

var (
	candidates []string               // Candidates stores each candidate for being the next word in the chain
	tracker    map[string]interface{} // Tracker keeps track of each valid word and the word before it
	count      int                    // Count keeps track of how many searches were done
)

var (
	ErrDifferentLengths    = fmt.Errorf("words must be of same length")
	ErrWordNotInDictionary = fmt.Errorf("word is not in the dictionary")
	ErrChainNotFound       = fmt.Errorf("couldn't find a valid chain")
	ErrSearchLimitReached  = fmt.Errorf("exceeded search limit")
)

func Execute(start, end string) ([]string, error) {
	// start word and end word can't be of different lengths
	if len(start) != len(end) {
		return nil, ErrDifferentLengths
	}

	// if the start word equals the end word we are done
	if start == end {
		return []string{start}, nil
	}

	dictionary, err := buildDictionary(len(start))
	if err != nil {
		return nil, fmt.Errorf("error building dictionary: %v", err)
	}

	// Check if the start word exists in the dictionary
	if err := checkExists(dictionary, start); err != nil {
		return nil, fmt.Errorf("%v: %s", err, start)
	}

	// Check if the end word exists in the dictionary
	if err := checkExists(dictionary, end); err != nil {
		return nil, fmt.Errorf("%v: %s", err, end)
	}

	// Initialize trackers
	candidates = append([]string{}, start)
	tracker = make(map[string]interface{})
	tracker[start] = nil
	count = 0

	return buildChain(end, dictionary)
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
			if !seen && isValid(candidate, word) {
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

func isValid(a, b string) bool {
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

// Opens the file containing the dictionary and builds
// the dictionary with words matching a certain length
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
