package main

import (
	"bufio"
	"fmt"
	"github.com/beevik/prefixtree"
	"os"
	"sort"
	"strings"
)

var allowedCharacters = []string{"E", "A", "G", "C", "N", "I", "F"}
var requiredCharacter = "F"

func main() {
	wordList, _ := readLines("Collins Scrabble Words (2019).txt")

	trie := prefixtree.New()

	for _, word := range wordList {
		trie.Add(word, true)
	}

	bag := make([]string, 0)

	recur([]string{}, trie, &bag)

	sort.Slice(bag, func(i, j int) bool {
		return len(bag[i]) > len(bag[j])
	})

	fmt.Println("All Matches: ")
	fmt.Println(bag)

	fmt.Println("\nAll Pentagrams:")

	for _, word := range bag {
		matched := true
		for _, character := range allowedCharacters {
			if !strings.Contains(word, character) {
				matched = false
				break
			}
		}

		if matched {
			fmt.Println(word)
		}
	}
}

func recur(letters []string, trie *prefixtree.Tree, bag *[]string) {
	word := strings.Join(letters[:], "")
	val, err := trie.Find(word)

	if err == prefixtree.ErrPrefixNotFound {
		return
	}

	if (val != nil || err != prefixtree.ErrPrefixAmbiguous) && strings.Contains(word, requiredCharacter) && len(word) > 3 {
		*bag = append(*bag, word)
	}

	for _, character := range allowedCharacters {
		recur(append(letters, character), trie, bag)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
