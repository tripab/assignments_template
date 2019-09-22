package cos418_hw1_1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Find the top K most common words in a text document.
// 	path: location of the document
//	numWords: number of words to return (i.e. k)
//	charThreshold: character threshold for whether a token qualifies as a word,
//		e.g. charThreshold = 5 means "apple" is a word but "pear" is not.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the same word.
// A word comprises alphanumeric characters only. All punctuations and other characters
// are removed, e.g. "don't" becomes "dont".
// You should use `checkError` to handle potential errors.
func topWords(path string, numWords int, charThreshold int) []WordCount {
	// TODO: implement me
	// HINT: You may find the `strings.Fields` and `strings.ToLower` functions helpful
	// HINT: To keep only alphanumeric characters, use the regex "[^0-9a-zA-Z]+"
	wc := make(map[string]int)
	reg := regexp.MustCompile("[^0-9a-zA-Z]+")

	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("Encountered error while opening the file %v! %v", path, err))
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for {
		if scanner.Scan() == false {
			break
		}
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			if len(word) >= charThreshold {
				wc[strings.ToLower(reg.ReplaceAllString(word, ""))]++
			}
		}
	}
	if scanner.Err() != nil {
		panic(fmt.Sprintf("Encountered error while reading the file %v! %v", path, scanner.Err()))
	}
	file.Close()

	return sortAndFetchTopN(wc, numWords)
}

func sortAndFetchTopN(wc map[string]int, numWords int) []WordCount {
	var topWordCounts []WordCount
	for k, v := range wc {
		topWordCounts = append(topWordCounts, WordCount{k, v})
	}
	sort.Slice(topWordCounts, func(i, j int) bool {
		if topWordCounts[i].Count == topWordCounts[j].Count {
			return topWordCounts[i].Word < topWordCounts[j].Word
		}
		return topWordCounts[i].Count > topWordCounts[j].Count
	})

	return topWordCounts[:numWords]
}

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
// DO NOT MODIFY THIS FUNCTION!
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}
