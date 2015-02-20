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
				fmt.Println(s, f)
                printFinger(i)
                str += string(s)
                fingersUsed = append(fingersUsed, i)
			}
		}
	}
    fmt.Println(str)
    fmt.Println(fingersUsed)

    //TODO make this separate function
    for i, d := range fingers {
        
    }
}

func main() {

	var input string
	fmt.Scanln(&input)
	getPossibleWords(input)

}
