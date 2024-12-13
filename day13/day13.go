package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Prize struct {
	X int
	Y int
}

type Button struct {
	X     int
	Y     int
	Cost  int
	Count int
}

type Machine struct {
	A     Button
	B     Button
	Prize Prize
}

func main() {

	file, err := os.Open("day13/day13.input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	machineList := []Machine{}
	machine := Machine{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "A") {
			xValue, err := strconv.Atoi(line[12:14])
			if err != nil {
				panic(err)
			}
			yValue, err := strconv.Atoi(line[18:20])
			if err != nil {
				panic(err)
			}
			machine.A = Button{Cost: 3, X: xValue, Y: yValue, Count: 0}
		} else if strings.Contains(line, "B") {
			xValue, err := strconv.Atoi(line[12:14])
			if err != nil {
				panic(err)
			}
			yValue, err := strconv.Atoi(line[18:20])
			if err != nil {
				panic(err)
			}
			machine.B = Button{Cost: 1, X: xValue, Y: yValue, Count: 0}
		} else if strings.Contains(line, "Prize") {
			xValue, err := strconv.Atoi(line[strings.Index(line, "X=")+2 : strings.Index(line, ",")])
			if err != nil {
				panic(err)
			}
			yValue, err := strconv.Atoi(line[strings.Index(line, "Y=")+2:])
			if err != nil {
				panic(err)
			}
			machine.Prize = Prize{X: xValue + 10000000000000, Y: yValue + 10000000000000}
			// machine.Prize = Prize{X: xValue, Y: yValue}
			machineList = append(machineList, machine)
			machine = Machine{}
		}
	}
	sum := 0
	for _, v := range machineList {
		sum += Play(v)

	}
	fmt.Printf("Totale: %v\n", sum)
}

func Play(machine Machine) int {

	// for i := 0; i < 300; i++ {
	// 	currentX := 0
	// 	currentY := 0
	// 	machine.A.Count = 0
	// 	machine.B.Count = 0

	// 	for j := 0; currentX < machine.Prize.X && currentY < machine.Prize.Y && machine.A.Count <= 100 && machine.B.Count <= 100; j++ {
	// 		if j < i {
	// 			currentX += machine.B.X
	// 			currentY += machine.B.Y
	// 			machine.B.Count++
	// 		} else {
	// 			currentX += machine.A.X
	// 			currentY += machine.A.Y
	// 			machine.A.Count++
	// 		}
	// 	}
	// 	if currentX == machine.Prize.X && currentY == machine.Prize.Y {
	// 		fmt.Printf("Found Prize: %v, Count A: %v, B: %v\n", machine.Prize, machine.A.Count, machine.B.Count)
	// 		return machine.A.Count*machine.A.Cost + machine.B.Count*machine.B.Cost
	// 	}

	// 	i++
	// fmt.Printf("No Prize: %v, current: %v %v Count A: %v, B: %v, i: %v\n", machine.Prize, currentX, currentY, machine.A.Count, machine.B.Count, i)
	// }
	// return 0

	// nom, dem := machine.Prize.Y*machine.B.X-machine.Prize.X*machine.B.Y, machine.A.Y*machine.B.X-machine.A.X*machine.B.Y
	// a := nom / dem
	// bx2 := machine.Prize.X - a*machine.A.X
	// if nom%dem == 0 && bx2%2 == 0 {
	// 	fmt.Printf("Prize: %v, Count A: %v, B: %v\n", machine.Prize, a, (bx2 / machine.B.X))
	// 	return 3*a + (bx2 / machine.B.X)
	// }
	X, Y := machine.Prize.X, machine.Prize.Y
	ax, ay := machine.A.X, machine.A.Y
	bx, by := machine.B.X, machine.B.Y

	d := ax*by - ay*bx
	da := by*X - bx*Y
	db := ax*Y - ay*X
	if da%d == 0 && db%d == 0 {
		// Integer solutions exist
		return 3*(da/d) + db/d
	}
	return 0
}
