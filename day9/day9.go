package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {

	file, err := os.Open("day9/day9.input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	// listPart1 := []string{}
	// for i := 0; i < len(line); i++ {
	// 	value, err := strconv.Atoi(string(line[i]))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	for j := 0; j < value; j++ {
	// 		if i%2 == 0 {
	// 			listPart1 = append(listPart1, strconv.Itoa(i/2))
	// 		} else {
	// 			listPart1 = append(listPart1, ".")
	// 		}
	// 	}
	// }
	// // fmt.Printf("START newList: %v\n", newlist)
	// fmt.Printf("delta: %v\n", len(listPart1)-Count(&listPart1, "."))
	// for i := len(listPart1) - 1; i >= 0; i-- {
	// 	if i < len(listPart1)-Count(&listPart1, ".") {
	// 		println("BREAK!!!!")
	// 		break
	// 	}
	// 	temp := listPart1[i]

	// 	listPart1[slices.Index(listPart1, ".")] = temp
	// 	listPart1[i] = "."
	// 	// fmt.Printf("newList: %v\n", newlist)
	// }
	// // fmt.Printf("FINList: %v\n", concatted)

	// sum := 0
	// for i, v := range listPart1 {
	// 	if v == "." {
	// 		break
	// 	}
	// 	intValue, err := strconv.Atoi(v)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	sum += i * intValue
	// }
	// fmt.Printf("Totale: %v\n", sum)

	listPart2 := []string{}
	for i := 0; i < len(line); i++ {
		value, err := strconv.Atoi(string(line[i]))
		if err != nil {
			panic(err)
		}
		for j := 0; j < value; j++ {
			if i%2 == 0 {
				listPart2 = append(listPart2, strconv.Itoa(i/2))
			} else {
				listPart2 = append(listPart2, ".")
			}
		}
	}
	fmt.Printf("listPart2 START:%v\n", listPart2)

	for i := len(listPart2) - 1; i >= 0; {
		if listPart2[i] == "." {
			i--
			continue
		}
		chunkLength := CountConsecutiveReverse(&listPart2, listPart2[i], i)
		start := slices.Index(listPart2, ".")
		for j := start; j < i; {
			fmt.Printf("\rcurrent i: %v, j: %v", i, j)
			pointIndex := slices.Index(listPart2[j:], ".") + j
			pointCount := CountConsecutive(&listPart2, ".", j)
			if pointCount >= chunkLength {
				// fmt.Printf("listPart2 PRE : %v\n", listPart2)
				// fmt.Printf("part: %v, chunkLength: %v, pointCount: %v, pointIndex: %v, start: %v, i: %v\n", listPart2[i], chunkLength, pointCount, pointIndex, j, i)
				temp := listPart2[pointIndex : pointIndex+chunkLength]
				chunk := listPart2[i-(chunkLength-1) : i+1]
				if pointIndex+chunkLength < len(listPart2) {
					temp = slices.Concat(temp, listPart2[i+1:])
				}
				// fmt.Printf("temp: %v, chunk: %v\n", temp, chunk)
				listPart2 = slices.Concat(listPart2[:pointIndex], chunk, listPart2[pointIndex+chunkLength:i-(chunkLength-1)], temp)

				// fmt.Printf("listPart2 POST: %v\n\n", listPart2)

				break
			}
			j = pointCount + pointIndex
		}
		if chunkLength > 0 {
			i = i - chunkLength
		} else {
			i--
		}
	}

	sum := 0

	for i, v := range listPart2 {
		if v == "." {
			continue
		}
		intValue, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		sum += i * intValue
	}
	fmt.Printf("Totale: %v\n", sum)
}

func Count(arr *[]string, match string) int {
	count := 0
	for _, v := range *arr {
		if v == match {
			count++
		}

	}
	return count
}

func CountConsecutive(arr *[]string, match string, start int) int {
	count := 0
	for i := start; i < len((*arr)); i++ {
		if (*arr)[i] == match {
			count++
		} else {
			break
		}
	}
	return count
}
func CountConsecutiveReverse(arr *[]string, match string, start int) int {
	count := 0
	for i := start; i > 0; i-- {
		if (*arr)[i] == match {
			count++
		} else {
			break
		}
	}
	return count
}
