package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	fmt.Println("Rucksack problem")
	fmt.Println("reading file")
	lines := readInput()
	fmt.Println(lines)
	fmt.Println("Reading every third line")
	sum := 0
	lineGroups := get3LineGroups(lines)
	for _, lineGroup := range lineGroups {
		char := findCommonChar(lineGroup)
		sum += getValueForCharacter(char)
	}
	fmt.Println(sum)
	fmt.Println("Found somtething")
}
func findCommonChar(group []string) string {
	fmt.Println(group)

	firstChars := findDuplicateChars(group[0], group[1])
	secondChars := findDuplicateChars(group[1], group[2])
	for _, sf := range firstChars {
		for _, ss := range secondChars {
			if ss == sf {
				return ss
			}
		}
	}
	fmt.Println(firstChars)
	fmt.Println(secondChars)
	panic("ps")
}

func get3LineGroups(lines []string) [][]string {
	result := [][]string{}
	tmp := []string{}
	for i, s := range lines {
		i++
		fmt.Println("-")
		fmt.Println(i)
		tmp = append(tmp, s)
		if i%3 == 0 {

			fmt.Print("-----")
			fmt.Println(tmp)
			result = append(result, tmp)
			tmp = []string{}
		}
	}

	return result
}

func partOne() {
	fmt.Println("Rucksack problem")
	fmt.Println("reading file")
	lines := readInput()
	dupChars := []string{}
	fmt.Println("Splitting lines in half")
	for _, s := range lines {
		left, right := splitInHalf(s)
		duplicateCharacters := findDuplicateChars(left, right)
		if len(duplicateCharacters) != 0 {
			dupChars = append(dupChars, duplicateCharacters...)
		}

	}
	fmt.Println("Found these duplicate chars")
	fmt.Println(dupChars)

	fmt.Println("Building sum")
	sum := buildSum(dupChars)
	fmt.Print("Total sum is: ")
	fmt.Println(sum)
}

func findDuplicateChars(left string, right string) []string {
	result := []string{}
	for i := 0; i < len(left); i++ {
		cl := left[i]
		for j := 0; j < len(right); j++ {
			if right[j] == cl {
				result = append(result, string(cl))
			}
		}
	}
	return result
}

func buildSum(chars []string) int {
	sum := 0
	for _, s := range chars {
		sum += getValueForCharacter(s)
	}

	return sum
}

func getValueForCharacter(char string) int {

	switch char {
	case "a":
		return 1
	case "b":
		return 2
	case "c":
		return 3
	case "d":
		return 4
	case "e":
		return 5
	case "f":
		return 6
	case "g":
		return 7
	case "h":
		return 8
	case "i":
		return 9
	case "j":
		return 10
	case "k":
		return 11
	case "l":
		return 12
	case "m":
		return 13
	case "n":
		return 14
	case "o":
		return 15
	case "p":
		return 16
	case "q":
		return 17
	case "r":
		return 18
	case "s":
		return 19
	case "t":
		return 20
	case "u":
		return 21
	case "v":
		return 22
	case "w":
		return 23
	case "x":
		return 24
	case "y":
		return 25
	case "z":
		return 26
	case "A":
		return 27
	case "B":
		return 28
	case "C":
		return 29
	case "D":
		return 30
	case "E":
		return 31
	case "F":
		return 32
	case "G":
		return 33
	case "H":
		return 34
	case "I":
		return 35
	case "J":
		return 36
	case "K":
		return 37
	case "L":
		return 38
	case "M":
		return 39
	case "N":
		return 40
	case "O":
		return 41
	case "P":
		return 42
	case "Q":
		return 43
	case "R":
		return 44
	case "S":
		return 45
	case "T":
		return 46
	case "U":
		return 47
	case "V":
		return 48
	case "W":
		return 49
	case "X":
		return 50
	case "Y":
		return 51
	case "Z":
		return 52

	}
	return 1000000
}

func splitInHalf(input string) (left string, right string) {
	if len(input)%2 != 0 {
		fmt.Println(input)
		panic("Can not split in two even halves")
	}
	return input[0:(len(input) / 2)], input[(len(input) / 2):]
}

func readInput() []string {
	fileHandle, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can not read file")
		os.Exit(1)
	}
	reader := bufio.NewReader(fileHandle)
	line, err := reader.ReadString('\n')
	results := []string{}

	lineNumber := 0
	for line != "" {
		if err != nil {
			fmt.Println("It happend on line " + strconv.Itoa(lineNumber))
			fmt.Println(err.Error())
			os.Exit(1)
		}
		results = append(results, strings.Trim(line, "\n"))
		lineNumber++
		line, err = reader.ReadString('\n')
	}
	fmt.Println("Read " + strconv.Itoa(lineNumber) + " lines")
	return results
}
