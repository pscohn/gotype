package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

var (
	letters string = "ETAION SHRDLU"
	m       map[string]map[int][]string

	leftPinky   = []string{"a", "q", "z"}
	leftRing    = []string{"s", "w", "x"}
	leftMiddle  = []string{"e", "d", "c"}
	leftIndex   = []string{"f", "t", "g", "v", "b", "r"}
	rightIndex  = []string{"h", "y", "n", "j", "u", "m"}
	rightMiddle = []string{"i", "k"}
	rightRing   = []string{"o", "l"}
	rightPinky  = []string{"p"}

	fingers = [][]string{leftPinky, leftRing, leftMiddle, leftIndex,
		rightIndex, rightMiddle, rightRing, rightPinky}
)

type Finger struct {
	name string
	keys []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func in(s string, a []string) bool {
	for _, r := range a {
		if s == string(r) {
			return true
		}
	}
	return false
}

func hasChar(char string, s string) bool {
	for _, r := range s {
		if char == string(r) {
			return true
		}
	}
	return false
}

func perms(input [][]string) []string {
	if len(input) == 0 {
		return []string{""}
	}

	var result []string
	strings := perms(input[1:])
	for _, c := range input[0] {
		for _, s := range strings {
			result = append(result, c+s)
		}
	}
	return result
}

func hasVowel(word string) bool {
	if hasChar("a", word) || hasChar("e", word) ||
		hasChar("i", word) || hasChar("o", word) ||
		hasChar("u", word) || hasChar("y", word) {
		return true
	}
	return false
}

func getPossibleWords(input string) []string {

	var fingersUsed []int
	for _, s := range input {
		for i, f := range fingers {
			if in(string(s), f) {
				fingersUsed = append(fingersUsed, i)
			}
		}
	}

	var list [][]string
	for _, d := range fingersUsed {
		list = append(list, fingers[d])
	}

	permutations := perms(list)
	var purged []string
	for _, word := range permutations {
		if hasVowel(word) {
			purged = append(purged, word)
		}
	}
	return purged
}

func max(arr []int) int {
	max := 0
	index := 0
	for i, n := range arr {
		if n > max {
			max = n
			index = i
		}
	}
	return index
}

func findBestWord(permutations []string) string {
	// alternate approach:
	// read file contents
	// get all words
	// check count of each permutation in words
	// return word with highest count

	count := make([]int, len(permutations))

	for i, word := range permutations {
		for _, w := range m[strings.ToLower(string(word[0]))][len(word)] {
			if strings.ToLower(word) == strings.ToLower(w) {
				count[i]++
			}
		}
	}
	return permutations[max(count)]
}

func mapDictionary(path string) {
	contents, err := ioutil.ReadFile(path)
	check(err)
	words := strings.Fields(string(contents))
	m = make(map[string]map[int][]string)
	for _, word := range words {
		key := strings.ToLower(string(word[0]))
		if m[key] == nil {
			m[key] = make(map[int][]string)
		}
		if m[key][len(word)] == nil {
			m[key][len(word)] = make([]string, 100)
		}
		m[key][len(word)] = append(m[key][len(word)], word)
	}
}

func init() {
	_, filename, _, _ := runtime.Caller(1)
	dictPath := path.Join(path.Dir(filename), "dict.txt")
	go mapDictionary(dictPath)
}

func main() {
	for {
		fmt.Printf("-> ")
		in := bufio.NewReader(os.Stdin)
		input, _ := in.ReadString('\n')
		split := strings.Split(input, " ")
		for _, sp := range split {
			permutations := getPossibleWords(sp)
			word := findBestWord(permutations)
			fmt.Printf("%s ", word)
		}
		fmt.Printf("\n")
	}
}
