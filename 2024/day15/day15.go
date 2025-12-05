package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("day15/day15.test3")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	mappa := [][]rune{}
	mappa2 := [][]rune{}
	moves := ""
	var robot []int
	var robot2 []int
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			riga2 := ""
			for _, v := range line {
				if v == '#' {
					riga2 += "##"
				} else if v == '@' {
					riga2 += "@."
				} else if v == 'O' {
					riga2 += "[]"
				} else if v == '.' {
					riga2 += ".."
				}
			}
			mappa2 = append(mappa2, []rune(riga2))
			mappa = append(mappa, []rune(line))
			if strings.Contains(line, "@") {
				robot = []int{strings.Index(line, "@"), len(mappa) - 1}
			}
			if strings.Contains(line, "@") {
				robot2 = []int{strings.Index(line, "@") * 2, len(mappa2) - 1}
			}
		} else {
			moves += line
		}
	}
	println("mappa")
	for _, riga := range mappa2 {
		for _, v := range riga {
			fmt.Printf("%s", string(v))
		}
		println()
	}
	fmt.Printf("\n robot at: %v robot2 at: %v\n", robot, robot2)
	fmt.Println(moves)

	for _, v := range moves {
		_, robot = MoveTo(&mappa, robot, v)
	}

	for _, v := range moves {
		_, robot2 = MoveTo(&mappa2, robot2, v)
	}

	println("mappa finale:")
	sum := 0
	for y, riga := range mappa2 {
		for x, v := range riga {
			fmt.Printf("%s", string(v))
			if string(v) == "O" {
				sum += y*100 + x
			}
		}
		println()
	}
	fmt.Printf("Totale: %v\n", sum)
}

func MoveTo(mappa *[][]rune, position []int, direction rune) ([]int, []int) {
	dX := 0
	dY := 0
	isBoxLeft := false
	isBoxRight := false
	if direction == 'v' {
		dY = 1
	} else if direction == '>' {
		dX = 1
	} else if direction == '<' {
		dX = -1
	} else {
		dY = -1
	}
	newPostion := []int{position[0] + dX, position[1] + dY}

	if dY != 0 {
		if (*mappa)[position[1]][position[0]] == '[' {
			isBoxLeft = true
		} else if (*mappa)[position[1]][position[0]] == ']' {
			isBoxRight = true
		}
	}

	if (*mappa)[newPostion[1]][newPostion[0]] == '#' {
		// println("found wall")
		return []int{position[0] - dX, position[1] - dY}, position
	}
	if (*mappa)[newPostion[1]][newPostion[0]] == '.' {
		if isBoxLeft || isBoxRight {
			if isBoxLeft {
				if (*mappa)[newPostion[1]][newPostion[0]+1] == '.' {
					Replace(mappa, position, newPostion)
					Replace(mappa, []int{position[0] + 1, position[1]}, []int{newPostion[0] + 1, newPostion[1]})
					return position, newPostion
				} else if (*mappa)[newPostion[1]][newPostion[0]+1] == '#' {
					return []int{position[0] - dX, position[1] - dY}, position
				}
			}
			if isBoxRight {
				if (*mappa)[newPostion[1]][newPostion[0]-1] == '.' {
					Replace(mappa, position, newPostion)
					Replace(mappa, []int{position[0] - 1, position[1]}, []int{newPostion[0] - 1, newPostion[1]})
					return position, newPostion
				} else if (*mappa)[newPostion[1]][newPostion[0]-1] == '#' {
					return []int{position[0] - dX, position[1] - dY}, position
				}
			}
		} else {
			Replace(mappa, position, newPostion)
			return position, newPostion
		}
	}
	var sidePosition []int
	if isBoxLeft || isBoxRight {
		if isBoxLeft {
			sidePosition, _ = MoveTo(mappa, []int{newPostion[0] + 1, newPostion[1]}, direction)
		}
		if isBoxRight {
			sidePosition, _ = MoveTo(mappa, []int{newPostion[0] - 1, newPostion[1]}, direction)
		}
	}

	actualPosition, _ := MoveTo(mappa, newPostion, direction)
	fmt.Printf("pos: %v, new:%v, actual: %v, direction: %v\n", position, newPostion, actualPosition, string(direction))
	if actualPosition[0] == newPostion[0] && actualPosition[1] == newPostion[1] {
		if isBoxLeft || isBoxRight {
			return []int{position[0] - dX, position[1] - dY}, position
		} else {
			if isBoxLeft || isBoxRight {
				if sidePosition != nil {
					Replace(mappa, position, actualPosition)
					if isBoxLeft {
						Replace(mappa, []int{position[0] + 1, position[1]}, sidePosition)
					} else if isBoxRight {
						Replace(mappa, []int{position[0] - 1, position[1]}, sidePosition)
					}
				} else {
					return []int{position[0] - dX, position[1] - dY}, position
				}
			} else {
				Replace(mappa, position, actualPosition)
			}
			return position, newPostion
		}
	} else {
		return []int{position[0] - dX, position[1] - dY}, position
	}
}

func Replace(mappa *[][]rune, old []int, new []int) {
	temp := (*mappa)[old[1]][old[0]]
	(*mappa)[old[1]][old[0]] = (*mappa)[new[1]][new[0]]
	(*mappa)[new[1]][new[0]] = temp
}
