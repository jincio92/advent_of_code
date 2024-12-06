package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("day3/day3.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	fullLine := ""
	for scanner.Scan() {
		fullLine += scanner.Text()

	}

	viableLine := ""
	doReg, err := regexp.Compile(`do\(\)`)
	if err != nil {
		panic(err)
	}

	dontReg, err := regexp.Compile(`don\'t\(\)`)
	if err != nil {
		panic(err)
	}
	doSplit := doReg.Split(fullLine, -1)

	for _, v := range doSplit {
		dontSplit := dontReg.Split(v, -1)
		viableLine += dontSplit[0]
	}
	fmt.Printf("newLine: %v", viableLine)
	println()

	reg, err := regexp.Compile(`mul\(([0-9]{1,3},[0-9]{1,3})\)`)
	if err != nil {
		panic(err)
	}

	indexes := reg.FindAllStringSubmatch(fullLine, -1)
	for _, v := range indexes {
		arguments := strings.Split(v[1], ",")
		firstArg, err := strconv.Atoi(arguments[0])
		if err != nil {
			panic(err)
		}
		secondArg, err := strconv.Atoi(arguments[1])
		if err != nil {
			panic(err)
		}
		sum += firstArg * secondArg
	}
	fmt.Printf("Somma: %v", sum)
	println()

	indexes = reg.FindAllStringSubmatch(viableLine, -1)
	sum = 0
	for _, v := range indexes {
		arguments := strings.Split(v[1], ",")
		firstArg, err := strconv.Atoi(arguments[0])
		if err != nil {
			panic(err)
		}
		secondArg, err := strconv.Atoi(arguments[1])
		if err != nil {
			panic(err)
		}
		sum += firstArg * secondArg
	}
	fmt.Printf("Somma Do/Don't: %v", sum)
	println()

}
