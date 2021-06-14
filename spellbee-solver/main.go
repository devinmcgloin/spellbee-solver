package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/beevik/prefixtree"
)

var requiredCharacter string
var allowedCharacters string

func init() {
	flag.StringVar(&requiredCharacter, "required", "F", "required character")
	flag.StringVar(&allowedCharacters, "allowed", "EAGCNIF", "allowed character")
}

func main() {
	flag.Parse()
	fmt.Printf("Solving for allowed characters %s, and required character %s\n", allowedCharacters, requiredCharacter)

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
	bag = unique(bag)

	fmt.Printf("All Matches (%d):\n", len(bag))
	fmt.Println(bag)

	fmt.Println("\nAll Pangrams:")

	for _, word := range bag {
		matched := true
		for _, character := range strings.Split(allowedCharacters, "") {
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

	for _, character := range strings.Split(allowedCharacters, "") {
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

func unique(array []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range array {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
