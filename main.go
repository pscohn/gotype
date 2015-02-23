package main

import "fmt"

type Finger struct {
	name string
	keys []string
}

func in(s string, a []string) bool {
	for _, r := range a {
		if s == string(r) {
			return true
		}
	}
	return false
}

func printFinger(i int) {
    var finger string
    switch i {
    case 0:
        finger = "left pinky" 
    case 1:
        finger = "left ring"
    case 2:
        finger = "left middle"
    case 3:
        finger = "left index"
    case 4:
        finger = "right index" 
    case 5:
        finger = "right middle"
    case 6:
        finger = "right ring"
    case 7:
        finger = "right pinky"
}
    fmt.Println(finger)
}


func perms(input [][]string) []string {
    if len(input) == 0 {
        return []string{""}
    }

    var result []string
    strings := perms(input[1:])
    for _, c := range input[0] {
        for _, s := range strings {
            result = append(result, c + s)
        }
    }
    return result
}

func getPossibleWords(input string) {
	leftPinky := []string{"a", "q", "z"}
	leftRing := []string{"s", "w", "x"}
	leftMiddle := []string{"e", "d", "c"}
	leftIndex := []string{"f", "t", "g", "v", "b", "r"}
	rightIndex := []string{"h", "y", "n", "j", "u", "m"}
	rightMiddle := []string{"i", "k", ","}
	rightRing := []string{"o", "l", "."}
	rightPinky := []string{"p", ";", "/"}

	fingers := [][]string{leftPinky, leftRing, leftMiddle, leftIndex,
		rightIndex, rightMiddle, rightRing, rightPinky}

    str := ""
    var fingersUsed []int
	for _, s := range input {
		for i, f := range fingers {
			if in(string(s), f) {
                str += string(s)
                fingersUsed = append(fingersUsed, i)
			}
		}
	}

    var list [][]string
    for _, d := range fingersUsed {
        list = append(list, fingers[d])
    }
    
    var permutations []string
    permutations = perms(list)
    fmt.Println("permutations:", permutations)
}

func main() {

	var input string
	fmt.Scanln(&input)
	getPossibleWords(input)

}
