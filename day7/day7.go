package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("day7/day7.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	testResult := []int{}
	digit := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		colonIndex := strings.Index(line, ":")

		testResultNumber, err := strconv.Atoi(line[:colonIndex])
		if err != nil {
			panic(err)
		}
		testResult = append(testResult, testResultNumber)
		splits := strings.Split(line[colonIndex+2:], " ")
		arr := []int{}
		for _, v := range splits {
			numb, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			arr = append(arr, numb)
		}
		digit = append(digit, arr)
	}
	sum := 0
	for i, d := range digit {
		comboLength := int(math.Pow(float64(2), float64(len(d)-1)))
		comboList := []string{}

		// fmt.Printf("comboLength: %v\n", comboLength)
		for i := 0; i < comboLength; i++ {
			comboList = append(comboList, fmt.Sprintf("%0"+strconv.Itoa(len(d)-1)+"s", strconv.FormatInt(int64(i), 2)))
		}
		for _, comb := range comboList {
			result := Evaluate(d, comb)
			if testResult[i] == result {
				fmt.Printf("testResult: %v, result: %v, combo: %v\n", testResult[i], result, comb)
				fmt.Printf("Trovato!\n")
				sum += testResult[i]
				break
			}
		}
		// fmt.Printf("digits: %v, bits: %v\n", d, comboList)
	}
	fmt.Printf("newTotal: %v\n", sum)
	sum = 0
	for i, d := range digit {
		comboLength := int(math.Pow(float64(3), float64(len(d)-1)))

		for j := 0; j < comboLength; j++ {
			// fmt.Printf("start run %d\n", j)
			result := EvaluateV2(d, j, len(d)-2)
			if testResult[i] == result {
				fmt.Printf("testResult: %v, result: %v\n", testResult[i], result)
				fmt.Printf("Trovato!\n")
				sum += testResult[i]
				break
			}
			// fmt.Printf("result run %d, total: %v\n", j, result)

		}
	}
	fmt.Printf("total: %v\n", sum)

}

func Evaluate(values []int, pattern string) int {

	if len(values) == 1 {
		return values[0]
	}
	if string(pattern[len(pattern)-1]) == "0" {
		return Evaluate(values[:len(values)-1], pattern[:len(pattern)-1]) + values[len(values)-1]
	} else if pattern == " " {
		result, err := strconv.Atoi(strconv.Itoa(Evaluate(values[:len(values)-1], pattern[:len(pattern)-1])) + strconv.Itoa(values[len(values)-1]))
		if err != nil {
			panic(err)
		}
		return result
	} else {
		return Evaluate(values[:len(values)-1], pattern[:len(pattern)-1]) * values[len(values)-1]
	}
}

func EvaluateV2(values []int, totalCombo int, depth int) int {
	// fmt.Printf("values: %v, total: %v, depth: %v\n", values, totalCombo, depth)

	if len(values) == 1 {
		return values[0]
	}
	nextTotal := totalCombo - ((int(math.Pow(float64(3), float64(depth)))) * (totalCombo / (int(math.Pow(float64(3), float64(depth))))))
	if nextTotal < 0 {
		nextTotal = totalCombo
	}
	evaluated := EvaluateV2(values[:len(values)-1], nextTotal, depth-1)
	if totalCombo/(int(math.Pow(float64(3), float64(depth)))) == 0 {
		return evaluated + values[len(values)-1]
	} else if totalCombo/(int(math.Pow(float64(3), float64(depth)))) == 1 {
		result, err := strconv.Atoi(strconv.Itoa(evaluated) + strconv.Itoa(values[len(values)-1]))
		if err != nil {
			panic(err)
		}
		return result
	} else {
		return evaluated * values[len(values)-1]
	}
}

func Operation(a int, b int, isMult bool) int {
	if isMult {
		return a * b
	} else {
		return a + b
	}
}
