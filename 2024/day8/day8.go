package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("day8/day8.input")
	if err != nil {
		panic(err)
	}
	mappaAntenna := map[string][][2]int{}
	scanner := bufio.NewScanner(file)
	mappa := []string{}

	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			if string(line[i]) != "." {
				if _, ok := mappaAntenna[string(line[i])]; !ok {
					mappaAntenna[string(line[i])] = [][2]int{}
				}
				mappaAntenna[string(line[i])] = append(mappaAntenna[string(line[i])], [2]int{i, lineNumber})
			}
		}
		mappa = append(mappa, line)
		lineNumber++
	}
	// fmt.Printf("parsato: %v\n", mappaAntenna)
	// v := mappaAntenna["0"]
	for _, v := range mappaAntenna {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				loopN := 0
				for true {
					delta := [2]int{v[i][0] - v[j][0], v[i][1] - v[j][1]}
					// fmt.Printf("v1: %v, v2: %v, delta: %v\n", v[i], v[j], delta)
					newPlusX := v[i][0] + (delta[0] * loopN)
					newPlusY := v[i][1] + (delta[1] * loopN)

					plusOut := false
					if newPlusY < len(mappa) && newPlusY >= 0 && newPlusX < len(mappa[0]) && newPlusX >= 0 {
						// fmt.Printf("new Vector X: %d, Y: %d\n", newPlusX, newPlusY)
						lineToEdit := mappa[newPlusY]
						lineToEdit = lineToEdit[:newPlusX] + "#" + lineToEdit[(newPlusX)+1:]
						mappa[newPlusY] = lineToEdit
					} else {
						plusOut = true
					}
					newMinusX := v[i][0] - (delta[0] * loopN)
					newMinusY := v[i][1] - (delta[1] * loopN)
					minusOut := false
					if newMinusY < len(mappa) && newMinusY >= 0 && newMinusX < len(mappa[0]) && newMinusX >= 0 {
						// fmt.Printf("new Vector X: %d, Y: %d\n", newMinusX, newMinusY)
						lineToEdit := mappa[newMinusY]
						lineToEdit = lineToEdit[:newMinusX] + "#" + lineToEdit[(newMinusX)+1:]
						mappa[newMinusY] = lineToEdit
					} else {
						minusOut = true
					}
					if minusOut && plusOut {
						break
					}
					loopN++
				}
			}
		}
	}
	sum := 0
	for _, v := range mappa {
		sum += strings.Count(v, "#")
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("Totale: %v\n", sum)
}
