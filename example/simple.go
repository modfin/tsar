package main

import (
	"bytes"
	"fmt"
	"github.com/modfin/henry/slicez"
	"github.com/modfin/tsar"
	"github.com/modfin/tsar/query"
	"io/ioutil"
	"regexp"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func tokenizeText(text string) []string {
	var splitter = regexp.MustCompile("(\\s+)|([!-/:-@[-`{-~])")
	return slicez.Filter(slicez.Map(splitter.Split(text, -1), func(word string) string {
		return strings.ToLower(strings.TrimSpace(word))
	}), func(s string) bool { return len(s) > 0 })
}

func main() {

	fileBytes, err := ioutil.ReadFile("sample.txt")
	check(err)
	text := string(fileBytes)
	wordlist := tsar.NewEntryList()

	var offset uint32
	for _, line := range strings.Split(text, "\n") {
		words := tokenizeText(line)

		for _, word := range words {
			check(wordlist.Append(word, offset)) // Adding the word with a byte offset to the line containing the word.
		}
		offset += uint32(len(line) + 1)
	}

	index := wordlist.ToIndex()

	entries, err := index.Find("swiftly", tsar.MatchEqual)
	check(err)
	for _, entry := range entries {
		fmt.Println(entry.Key, "- exist on lines staring with the byte offsets", entry.Pointers)
	}

	fmt.Println()
	fmt.Println("Query: (king & scandi:*) | Threatens)")
	// Marshaling the index, it can be persisted on disk and using the file interface, read directly by query
	indexBytes := tsar.MarshalIndex(index)
	result, err := query.Query("(king & scandi:*) | beautiful", bytes.NewReader(fileBytes), bytes.NewReader(indexBytes), 5, 0)
	check(err)

	fmt.Println(string(result))

}
